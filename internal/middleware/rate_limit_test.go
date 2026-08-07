package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRateLimiterAllowsBurstThenBlocks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rl := NewRateLimiter(1, 3)
	router := gin.New()
	router.Use(rl.Middleware())
	router.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })

	statuses := make([]int, 0, 6)
	for i := 0; i < 6; i++ {
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		req.RemoteAddr = "203.0.113.1:1234"
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		statuses = append(statuses, rec.Code)
	}

	for i, code := range statuses {
		want := http.StatusOK
		if i >= 3 {
			want = http.StatusTooManyRequests
		}
		if code != want {
			t.Fatalf("request %d: got status %d, want %d", i+1, code, want)
		}
	}
}

func TestRateLimiterRecoversAfterDelay(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rl := NewRateLimiter(2, 1)
	router := gin.New()
	router.Use(rl.Middleware())
	router.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })

	url := "/ping"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.RemoteAddr = "198.51.100.7:5555"

	// First request succeeds.
	rec := httptest.NewRecorder()
	if router.ServeHTTP(rec, req); rec.Code != http.StatusOK {
		t.Fatalf("first request: got %d, want 200", rec.Code)
	}

	time.Sleep(700 * time.Millisecond)
	rec = httptest.NewRecorder()
	if router.ServeHTTP(rec, req); rec.Code != http.StatusOK {
		t.Fatalf("second request after wait: got %d, want 200", rec.Code)
	}
}

func TestRateLimiterIsolationPerClient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	rl := NewRateLimiter(1000, 5)
	router := gin.New()
	router.Use(rl.Middleware())
	router.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })

	serve := func(ip string) int {
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		req.RemoteAddr = ip + ":5000"
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		return rec.Code
	}

	// Client A exhausts its own budget.
	for i := 0; i < 5; i++ {
		if code := serve("192.0.2.1"); code != http.StatusOK {
			t.Fatalf("client A request %d: got %d, want 200", i+1, code)
		}
	}
	if code := serve("192.0.2.1"); code != http.StatusTooManyRequests {
		t.Fatalf("client A over budget: got %d, want 429", code)
	}

	// Client B is unaffected by A.
	for i := 0; i < 5; i++ {
		if code := serve("192.0.2.2"); code != http.StatusOK {
			t.Fatalf("client B request %d: got %d, want 200", i+1, code)
		}
	}
}
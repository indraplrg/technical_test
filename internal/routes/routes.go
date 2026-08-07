package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/indraplrg/technical_test/internal/config"
	"github.com/indraplrg/technical_test/internal/controller"
	"github.com/indraplrg/technical_test/internal/middleware"
	"github.com/indraplrg/technical_test/internal/repository"
	"github.com/indraplrg/technical_test/internal/response"
	"github.com/indraplrg/technical_test/internal/service"
)

// Setup wires all dependencies and returns the configured Gin engine.
func Setup(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), middleware.Recovery(), middleware.RequestID(), middleware.CORS(cfg.AllowedOrigin))

	router.GET("/health", health)

	v1 := router.Group("/api/v1")
	{
		jurusanController := newJurusanController(db)
		jurusanRoutes := v1.Group("/jurusan")
		{
			jurusanRoutes.GET("", jurusanController.GetAll)
			jurusanRoutes.GET("/:id", jurusanController.GetByID)
			jurusanRoutes.POST("", jurusanController.Create)
			jurusanRoutes.PUT("/:id", jurusanController.Update)
			jurusanRoutes.DELETE("/:id", jurusanController.Delete)
		}
	}

	return router
}

func health(c *gin.Context) {
	response.Success(c, http.StatusOK, "service healthy", gin.H{"status": "ok"})
}

func newJurusanController(db *gorm.DB) *controller.JurusanController {
	jurusanRepo := repository.NewJurusanRepository(db)
	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	jurusanService := service.NewJurusanService(jurusanRepo, mahasiswaRepo)
	return controller.NewJurusanController(jurusanService)
}

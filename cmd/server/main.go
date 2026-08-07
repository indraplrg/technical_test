package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/indraplrg/technical_test/internal/config"
	"github.com/indraplrg/technical_test/internal/database"
	"github.com/indraplrg/technical_test/internal/model"
	"github.com/indraplrg/technical_test/internal/routes"

	_ "github.com/indraplrg/technical_test/docs" // swagger generated docs
)

// @title Student Management API
// @version 1.0
// @description REST API for managing Jurusan (departments) and Mahasiswa (students), built with Go, Gin, GORM and PostgreSQL.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey
// @in header
// @name Authorization
// @description Insert the token into the Authorization header

func main() {
	migrateFlag := flag.Bool("migrate", false, "run auto migrations and seed, then exit")
	flag.Parse()

	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	slog.Info("database connection ready", "db", cfg.DBName)

	if err := database.RunMigrations(db, &model.Jurusan{}, &model.Mahasiswa{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	if err := database.Seed(db); err != nil {
		log.Fatalf("failed to seed data: %v", err)
	}

	if *migrateFlag {
		slog.Info("migrations and seed completed")
		return
	}

	router := routes.Setup(cfg, db)

	server := &http.Server{
		Addr:           ":" + cfg.AppPort,
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(cfg.IdleTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		slog.Info("server started", "port", cfg.AppPort, "env", cfg.AppEnv)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	slog.Info("server stopped cleanly")
}

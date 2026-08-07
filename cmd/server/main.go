package main

import (
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/config"
	"github.com/indraplrg/technical_test/internal/database"
	"github.com/indraplrg/technical_test/internal/model"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	slog.Info("database connection ready", "db", cfg.DBName)

	if err := database.RunMigrations(db, &model.Jurusan{}, &model.Mahasiswa{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "ok"})
	})

	log.Fatal(router.Run(":" + cfg.AppPort))
}
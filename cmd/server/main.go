package main

import (
	"log/slog"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/config"
	"github.com/indraplrg/technical_test/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	slog.Info("database connection ready", "db", cfg.DBName)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "ok"})
	})

	_ = db

	log.Fatal(router.Run(":" + cfg.AppPort))
}
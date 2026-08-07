package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/config"
)

func main() {
	cfg := config.Load()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "ok"})
	})

	log.Fatal(router.Run(":" + cfg.AppPort))
}
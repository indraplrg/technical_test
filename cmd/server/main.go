package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "ok"})
	})

	log.Fatal(router.Run())
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"sectrack/config"
	"sectrack/models"
	"sectrack/routes"
)
func main() {
	godotenv.Load()

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Vulnerability{},
		&models.AuditLog{},
	)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SecTrack API",
			"version": "1.0.0",
			"status":  "running",
		})
	})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
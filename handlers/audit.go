package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sectrack/config"
	"sectrack/models"
)

func GetAuditLogs(c *gin.Context) {
	var logs []models.AuditLog

	config.DB.Order("created_at desc").Find(&logs)

	c.JSON(http.StatusOK, gin.H{"data": logs})
}
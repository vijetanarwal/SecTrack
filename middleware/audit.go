package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sectrack/config"
	"sectrack/models"
)

func AuditLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		method := c.Request.Method
		if method == "POST" || method == "PATCH" || method == "DELETE" {
			userID, exists := c.Get("user_id")
			if !exists {
				return
			}

			action := fmt.Sprintf("%s %s", method, c.FullPath())

			log := models.AuditLog{
				UserID:   userID.(uint),
				Action:   action,
				Resource: c.FullPath(),
			}

			config.DB.Create(&log)
		}
	}
}
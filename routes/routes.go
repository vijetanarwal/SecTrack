package routes

import (
	"github.com/gin-gonic/gin"
	"sectrack/handlers"
	"sectrack/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth)
	protected.Use(middleware.AuditLogger())
	{
		protected.POST("/vulns", middleware.RequireRole("admin", "analyst"), handlers.CreateVulnerability)
		protected.GET("/vulns", handlers.GetVulnerabilities)
		protected.GET("/vulns/:id", handlers.GetVulnerability)
		protected.PATCH("/vulns/:id", middleware.RequireRole("admin", "analyst"), handlers.UpdateVulnerability)
		protected.DELETE("/vulns/:id", middleware.RequireRole("admin"), handlers.DeleteVulnerability)

		protected.GET("/audit-logs", middleware.RequireRole("admin"), handlers.GetAuditLogs)
	}
}
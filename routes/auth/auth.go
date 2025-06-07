package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/payslip/config"
	"github.com/payslip/controllers"
	"github.com/payslip/middlewares"
	"github.com/payslip/services"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, s *services.Service) {
	auth_group := r.Group("/auth")
	auth_group.Use(middlewares.AuditMiddleware(s))
	auth_group.POST("admin/login", func(c *gin.Context) { controllers.AdminLogin(c, s) })
	auth_group.POST("employee/login", func(c *gin.Context) { controllers.EmployeeLogin(c, s) })

}

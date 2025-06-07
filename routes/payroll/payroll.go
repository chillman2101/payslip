package payroll

import (
	"github.com/gin-gonic/gin"
	"github.com/payslip/config"
	"github.com/payslip/controllers"
	"github.com/payslip/middlewares"
	"github.com/payslip/services"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, s *services.Service) {
	payroll_group := r.Group("admin/payroll")
	payroll_group.Use(middlewares.AuthMiddleware("employee"))
	payroll_group.Use(middlewares.AuditMiddleware(s))
	payroll_group.POST("/add", func(c *gin.Context) { controllers.AddPayroll(c, s) })

}

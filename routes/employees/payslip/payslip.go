package payslip

import (
	"github.com/gin-gonic/gin"
	"github.com/payslip/config"
	"github.com/payslip/controllers"
	"github.com/payslip/middlewares"
	"github.com/payslip/services"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, s *services.Service) {
	payslip_employee_group := r.Group("employee/payslip")
	payslip_employee_group.Use(middlewares.AuthMiddleware("employee"))
	payslip_employee_group.Use(middlewares.AuditMiddleware(s))
	payslip_employee_group.GET("/generate/:id", func(c *gin.Context) { controllers.SummaryPayroll(c, s) })

	payslip_admin_group := r.Group("admin/payslip")
	payslip_admin_group.Use(middlewares.AuthMiddleware("admin"))
	payslip_admin_group.Use(middlewares.AuditMiddleware(s))
	payslip_admin_group.GET("/generate/:id", func(c *gin.Context) { controllers.SummaryPayrollAdmin(c, s) })
}

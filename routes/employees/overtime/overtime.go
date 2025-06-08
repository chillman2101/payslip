package overtime

import (
	"github.com/gin-gonic/gin"
	"github.com/payslip/config"
	"github.com/payslip/controllers"
	"github.com/payslip/middlewares"
	"github.com/payslip/services"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, s *services.Service) {
	overtime_group := r.Group("employee/overtime")
	overtime_group.Use(middlewares.AuthMiddleware("employee"))
	overtime_group.Use(middlewares.AuditMiddleware(s))
	overtime_group.Use(middlewares.PayrollMiddleware(s))
	overtime_group.POST("/submit", func(c *gin.Context) { controllers.SubmitOvertime(c, s) })
}

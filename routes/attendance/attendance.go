package attendance

import (
	"github.com/gin-gonic/gin"
	"github.com/payslip/config"
	"github.com/payslip/controllers"
	"github.com/payslip/middlewares"
	"github.com/payslip/services"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, s *services.Service) {
	attendance_group := r.Group("employee/attendance")
	attendance_group.Use(middlewares.AuthMiddleware("employee"))
	attendance_group.Use(middlewares.AuditMiddleware(s))
	attendance_group.POST("/check-in", func(c *gin.Context) { controllers.CheckIn(c, s) })
	attendance_group.POST("/check-out", func(c *gin.Context) { controllers.CheckOut(c, s) })

}

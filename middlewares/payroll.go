package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/payslip/models"
	"github.com/payslip/services"
)

func PayrollMiddleware(s *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // fake
		payroll_req := models.Payroll{
			StartDate: start,
			EndDate:   start,
		}

		payroll, err := s.PayrollRepo.FindPayrollByDate(context.Background(), &payroll_req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error Check Payroll " + err.Error()})
			return
		}

		// cannot create attendance, overtime, and reimbursement if payroll is not found
		if payroll == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Payroll not found"})
			return
		}

		// cannot create attendance, overtime, and reimbursement if payroll was already proceed
		if payroll.AlreadyProceed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Payroll already proceed "})
			return
		}

		c.Set("payroll_id", payroll.ID)
		c.Next()
	}
}

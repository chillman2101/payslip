package integration_test

import (
	"context"
	"testing"

	"github.com/payslip/models"
	"github.com/payslip/repositories"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
)

func TestAdminPayroll_Integration(t *testing.T) {
	db := setupTestDB(t)
	tx := db

	payrollRepo := repositories.NewPayrollRepository(tx)
	svc := services.NewService(tx, nil, nil, payrollRepo, nil, nil, nil)

	req := models.AddPayrollRequest{
		Description: "Test Payroll Bulan Juni",
		StartDate:   "2025-06-01",
		EndDate:     "2025-06-30",
	}
	ctx := context.Background()
	_, err := svc.AddPayroll(ctx, req)
	assert.NoError(t, err, "checkout succeed")
}

package integration_test

import (
	"context"
	"testing"

	"github.com/payslip/models"
	"github.com/payslip/repositories"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
)

func TestSubmitOvertime_Success_Integration(t *testing.T) {
	db := setupTestDB(t)
	tx := db

	overtimeRepo := repositories.NewOvertimeRepository(tx)
	svc := services.NewService(tx, nil, nil, nil, nil, overtimeRepo, nil)

	employeeId := uint(1)
	payrollId := uint(1)

	ctx := context.Background()
	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId:  employeeId,
		PayrollId:   payrollId,
		AmountTime:  1,
		Description: "Test overtime",
	}

	_, err := svc.SubmitOvertime(ctx, req)
	assert.NoError(t, err, "check-in should succeed")

}

func TestSubmitOvertime_Over3Hours_Integration(t *testing.T) {
	db := setupTestDB(t)
	tx := db

	overtimeRepo := repositories.NewOvertimeRepository(tx)
	svc := services.NewService(tx, nil, nil, nil, nil, overtimeRepo, nil)

	employeeId := uint(1)
	payrollId := uint(1)

	ctx := context.Background()
	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId:  employeeId,
		PayrollId:   payrollId,
		AmountTime:  3,
		Description: "Test overtime",
	}

	_, err := svc.SubmitOvertime(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "Overtime cannot be requested up to 3 hours", err.Error())
}

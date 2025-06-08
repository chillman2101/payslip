package unit_test

import (
	"context"
	"testing"

	"github.com/payslip/models"
	"github.com/payslip/repositories/mocks"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddPayroll_Success(t *testing.T) {
	ctx := context.Background()

	mockPayrollRepo := new(mocks.PayrollRepository)
	svc := services.NewService(nil, nil, nil, mockPayrollRepo, nil, nil, nil)

	req := models.AddPayrollRequest{
		StartDate:   "2025-06-01",
		EndDate:     "2025-06-30",
		Description: "Payroll Bulan Juni",
	}

	// Parse date seperti di service
	// startDate, _ := time.Parse("2006-01-02", req.StartDate)
	// endDate, _ := time.Parse("2006-01-02", req.EndDate)

	mockPayrollRepo.On("FindPayrollByDate", ctx, mock.AnythingOfType("*models.Payroll")).
		Return(nil, nil).Once()

	mockPayrollRepo.On("CreatePayroll", ctx, mock.AnythingOfType("*models.Payroll")).
		Return(nil).Once()

	resp, err := svc.AddPayroll(ctx, req)

	assert.NoError(t, err)
	assert.Nil(t, resp)

	mockPayrollRepo.AssertExpectations(t)
}

func TestAddPayroll_InvalidDate(t *testing.T) {
	ctx := context.Background()
	svc := services.NewService(nil, nil, nil, nil, nil, nil, nil)

	req := models.AddPayrollRequest{
		StartDate:   "2025-13-01", // invalid month
		EndDate:     "2025-06-30",
		Description: "Invalid Date",
	}

	resp, err := svc.AddPayroll(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot parse")
}

func TestAddPayroll_Duplicate(t *testing.T) {
	ctx := context.Background()

	mockPayrollRepo := new(mocks.PayrollRepository)
	svc := services.NewService(nil, nil, nil, mockPayrollRepo, nil, nil, nil)

	req := models.AddPayrollRequest{
		StartDate:   "2025-06-01",
		EndDate:     "2025-06-30",
		Description: "Duplicate Payroll",
	}

	mockPayrollRepo.On("FindPayrollByDate", ctx, mock.AnythingOfType("*models.Payroll")).
		Return(&models.Payroll{Model: gorm.Model{ID: 1}}, nil).Once()

	resp, err := svc.AddPayroll(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "payroll already exist in this range of date", err.Error())

	mockPayrollRepo.AssertExpectations(t)
}

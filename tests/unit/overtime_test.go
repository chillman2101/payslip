package unit_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/payslip/models"
	"github.com/payslip/repositories/mocks"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestSubmitOvertime_NewOvertimeSuccess(t *testing.T) {
	ctx := context.Background()
	mockOvertimeRepo := new(mocks.OvertimeRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, mockOvertimeRepo, nil)

	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId: 1,
		AmountTime: 2,
		PayrollId:  1,
	}

	mockOvertimeRepo.On("FindOvertimeByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(nil, nil).Once()

	mockOvertimeRepo.On("InsertOvertime", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(nil, nil).Once()

	resp, err := svc.SubmitOvertime(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, resp)
	mockOvertimeRepo.AssertExpectations(t)
}

func TestSubmitOvertime_AccumulatedOvertimeSuccess(t *testing.T) {
	ctx := context.Background()
	mockOvertimeRepo := new(mocks.OvertimeRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, mockOvertimeRepo, nil)

	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId: 1,
		AmountTime: 1,
		PayrollId:  1,
	}

	existing := &models.Overtime{
		Model:      gorm.Model{ID: 100, CreatedAt: time.Now().Add(-time.Hour)},
		EmployeeId: 1,
		TotalHour:  1,
	}

	mockOvertimeRepo.On("FindOvertimeByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(existing, nil).Once()

	mockOvertimeRepo.On("InsertOvertime", ctx, mock.MatchedBy(func(o *models.Overtime) bool {
		return o.ID == 100 && o.TotalHour == 2
	})).Return(nil, nil).Once()

	resp, err := svc.SubmitOvertime(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, resp)
	mockOvertimeRepo.AssertExpectations(t)
}

func TestSubmitOvertime_OverLimitError(t *testing.T) {
	ctx := context.Background()
	mockOvertimeRepo := new(mocks.OvertimeRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, mockOvertimeRepo, nil)

	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId: 1,
		AmountTime: 2,
		PayrollId:  1,
	}

	existing := &models.Overtime{
		Model:      gorm.Model{ID: 100, CreatedAt: time.Now().Add(-time.Hour)},
		EmployeeId: 1,
		TotalHour:  2,
	}

	mockOvertimeRepo.On("FindOvertimeByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(existing, nil).Once()

	resp, err := svc.SubmitOvertime(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "Overtime cannot be requested up to 3 hours", err.Error())

	mockOvertimeRepo.AssertExpectations(t)
}

func TestSubmitOvertime_FindError(t *testing.T) {
	ctx := context.Background()
	mockOvertimeRepo := new(mocks.OvertimeRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, mockOvertimeRepo, nil)

	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId: 1,
		AmountTime: 2,
		PayrollId:  1,
	}

	mockOvertimeRepo.On("FindOvertimeByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(nil, errors.New("find error")).Once()

	resp, err := svc.SubmitOvertime(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "find error", err.Error())

	mockOvertimeRepo.AssertExpectations(t)
}

func TestSubmitOvertime_InsertError(t *testing.T) {
	ctx := context.Background()
	mockOvertimeRepo := new(mocks.OvertimeRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, mockOvertimeRepo, nil)

	req := models.EmployeeSubmitOvertimeRequest{
		EmployeeId: 1,
		AmountTime: 2,
		PayrollId:  1,
	}

	mockOvertimeRepo.On("FindOvertimeByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(nil, nil).Once()

	mockOvertimeRepo.On("InsertOvertime", ctx, mock.AnythingOfType("*models.Overtime")).
		Return(nil, errors.New("insert error")).Once()

	resp, err := svc.SubmitOvertime(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())

	mockOvertimeRepo.AssertExpectations(t)
}

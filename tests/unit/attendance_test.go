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
)

func TestCheckIn_Success(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
		PayrollId:  10,
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, nil).Once()

	mockAttendanceRepo.On("CheckInAttendance", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, nil).Once()

	resp, err := svc.CheckIn(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, resp)
	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckIn_AlreadyCheckedIn(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
		PayrollId:  10,
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(&models.Attendance{EmployeeId: req.EmployeeId}, nil).Once()

	resp, err := svc.CheckIn(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "employee has already checked in", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckIn_FindError(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
		PayrollId:  10,
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, errors.New("database error")).Once()

	resp, err := svc.CheckIn(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckIn_CheckInError(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
		PayrollId:  10,
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, nil).Once()

	mockAttendanceRepo.On("CheckInAttendance", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, errors.New("insert error")).Once()

	resp, err := svc.CheckIn(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckOut_Success(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
	}

	existingAttendance := &models.Attendance{
		EmployeeId:     req.EmployeeId,
		AttendanceDate: time.Now().Format("2006-01-02"),
		CheckOutTime:   time.Time{}, // zero value, not checked out yet
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(existingAttendance, nil).Once()

	mockAttendanceRepo.On("CheckOutAttendance", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, nil).Once()

	resp, err := svc.CheckOut(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, resp)
	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckOut_NotCheckedIn(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, nil).Once()

	resp, err := svc.CheckOut(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "employee has not checked in", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckOut_AlreadyCheckedOut(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
	}

	existingAttendance := &models.Attendance{
		EmployeeId:     req.EmployeeId,
		AttendanceDate: time.Now().Format("2006-01-02"),
		CheckOutTime:   time.Now(), // already checked out
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(existingAttendance, nil).Once()

	resp, err := svc.CheckOut(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "employee has already checked out", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckOut_FindError(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, errors.New("db error")).Once()

	resp, err := svc.CheckOut(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

func TestCheckOut_CheckOutError(t *testing.T) {
	ctx := context.Background()
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	svc := services.NewService(nil, nil, nil, nil, mockAttendanceRepo, nil, nil)

	req := models.EmployeeAttendanceRequest{
		EmployeeId: 1,
	}

	existingAttendance := &models.Attendance{
		EmployeeId:     req.EmployeeId,
		AttendanceDate: time.Now().Format("2006-01-02"),
		CheckOutTime:   time.Time{}, // not checked out yet
	}

	mockAttendanceRepo.On("FindAttendanceByEmployeeIdAndDate", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(existingAttendance, nil).Once()

	mockAttendanceRepo.On("CheckOutAttendance", ctx, mock.AnythingOfType("*models.Attendance")).
		Return(nil, errors.New("update error")).Once()

	resp, err := svc.CheckOut(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "update error", err.Error())

	mockAttendanceRepo.AssertExpectations(t)
}

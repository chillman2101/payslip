package unit_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/payslip/models"
	"github.com/payslip/repositories/mocks"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLoginEmployee_Success(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.AuthRepository)
	svc := services.NewService(nil, nil, mockRepo, nil, nil, nil, nil)

	expected_employee := &models.Employee{
		Model: gorm.Model{
			ID: 1,
		},
		Username: "bmabe0",
		Password: "$2a$05$KPn1xtrVrF/rJLRnox4PyuKp4QL0WOs34J6iBPHS/lLtT7NKrZCKO",
		Salary:   2782000,
	}

	mockRepo.On("FindEmployeeByUsername", ctx, "bmabe0").Return(expected_employee, nil)

	// call login
	reqID := uuid.NewString()
	req := models.UserRequest{
		RequestId: reqID,
		Username:  "bmabe0",
		Password:  "bmabe0",
	}
	result, err := svc.LoginEmployee(ctx, req)

	var employee models.Employee
	resp, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(resp, &employee)

	// verifikasi
	assert.NoError(t, err)
	assert.Equal(t, expected_employee.Username, employee.Username)
	mockRepo.AssertExpectations(t)
}

func TestLoginEmployee_WrongPassword(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.AuthRepository)
	svc := services.NewService(nil, nil, mockRepo, nil, nil, nil, nil)

	expected_employee := &models.Employee{
		Model: gorm.Model{
			ID: 1,
		},
		Username: "bmabe0",
		Password: "$2a$05$KPn1xtrVrF/rJLRnox4PyuKp4QL0WOs34J6iBPHS/lLtT7NKrZCKO",
		Salary:   2782000,
	}

	mockRepo.On("FindEmployeeByUsername", ctx, "bmabe0").Return(expected_employee, nil)

	// call login
	reqID := uuid.NewString()
	req := models.UserRequest{
		RequestId: reqID,
		Username:  "bmabe0",
		Password:  "admin123",
	}
	result, err := svc.LoginEmployee(ctx, req)

	var employee models.Employee
	resp, _ := json.Marshal(result)
	json.Unmarshal(resp, &employee)

	// verifikasi
	assert.Error(t, err)
	assert.Equal(t, "invalid password", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestLoginEmployee_NotFound(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.AuthRepository)
	svc := services.NewService(nil, nil, mockRepo, nil, nil, nil, nil)

	mockRepo.On("FindEmployeeByUsername", ctx, "bmabe023").Return(nil, errors.New("employee not found"))

	// call login
	reqID := uuid.NewString()
	req := models.UserRequest{
		RequestId: reqID,
		Username:  "bmabe023",
		Password:  "bmabe0",
	}
	_, err := svc.LoginEmployee(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "employee not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin_Success(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.AuthRepository)
	svc := services.NewService(nil, nil, mockRepo, nil, nil, nil, nil)

	expected_admin := &models.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		Username: "admin",
		Password: "$2a$05$HOLnXWH.HXZTyJC.bJy8TeuTA67dWyYtOZiTiNBXKYcitYTvwaDiK",
	}

	mockRepo.On("FindAdminByUsername", ctx, "admin").Return(expected_admin, nil)

	// call login
	reqID := uuid.NewString()
	req := models.UserRequest{
		RequestId: reqID,
		Username:  "admin",
		Password:  "admin",
	}
	result, err := svc.LoginAdmin(ctx, req)

	var admin models.Admin
	resp, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(resp, &admin)

	// verifikasi
	assert.NoError(t, err)
	assert.Equal(t, expected_admin.Username, admin.Username)
	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin_WrongPassword(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.AuthRepository)
	svc := services.NewService(nil, nil, mockRepo, nil, nil, nil, nil)

	expected_admin := &models.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		Username: "admin",
		Password: "$2a$05$KPn1xtrVrF/rJLRnox4PyuKp4QL0WOs34J6iBPHS/lLtT7NKrZCKO",
	}

	mockRepo.On("FindAdminByUsername", ctx, "admin").Return(expected_admin, nil)

	// call login
	reqID := uuid.NewString()
	req := models.UserRequest{
		RequestId: reqID,
		Username:  "admin",
		Password:  "admin123",
	}
	result, err := svc.LoginAdmin(ctx, req)

	var admin models.Admin
	resp, _ := json.Marshal(result)
	json.Unmarshal(resp, &admin)

	// verifikasi
	assert.Error(t, err)
	assert.Equal(t, "invalid password", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin_NotFound(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.AuthRepository)
	svc := services.NewService(nil, nil, mockRepo, nil, nil, nil, nil)

	mockRepo.On("FindAdminByUsername", ctx, "admin1234").Return(nil, errors.New("admin not found"))

	// call login
	reqID := uuid.NewString()
	req := models.UserRequest{
		RequestId: reqID,
		Username:  "admin1234",
		Password:  "admin",
	}
	_, err := svc.LoginAdmin(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "admin not found", err.Error())
	mockRepo.AssertExpectations(t)
}

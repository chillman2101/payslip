package unit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/payslip/models"
	"github.com/payslip/repositories/mocks"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmitReimbursement_Success(t *testing.T) {
	ctx := context.Background()
	mockReimbursementRepo := new(mocks.ReimbursementRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, nil, mockReimbursementRepo)

	req := models.EmployeeSubmitReimbursementRequest{
		EmployeeId:  1,
		Amount:      100000,
		Description: "Medical",
		PayrollId:   1,
	}

	mockReimbursementRepo.On("InsertReimbursement", ctx, mock.MatchedBy(func(r *models.Reimbursement) bool {
		return r.EmployeeId == req.EmployeeId &&
			r.TotalAmount == req.Amount &&
			r.Description == req.Description &&
			r.PayrollId == req.PayrollId
	})).Return(nil, nil).Once()

	resp, err := svc.SubmitReimbursement(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, resp)
	mockReimbursementRepo.AssertExpectations(t)
}

func TestSubmitReimbursement_InsertError(t *testing.T) {
	ctx := context.Background()
	mockReimbursementRepo := new(mocks.ReimbursementRepository)
	svc := services.NewService(nil, nil, nil, nil, nil, nil, mockReimbursementRepo)

	req := models.EmployeeSubmitReimbursementRequest{
		EmployeeId:  1,
		Amount:      100000,
		Description: "Transport",
		PayrollId:   2,
	}

	mockReimbursementRepo.On("InsertReimbursement", ctx, mock.AnythingOfType("*models.Reimbursement")).
		Return(nil, errors.New("insert error")).Once()

	resp, err := svc.SubmitReimbursement(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())

	mockReimbursementRepo.AssertExpectations(t)
}

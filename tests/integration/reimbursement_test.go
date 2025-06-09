package integration_test

import (
	"context"
	"testing"

	"github.com/payslip/models"
	"github.com/payslip/repositories"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
)

func TestSubmitReimbursement_Success_Integration(t *testing.T) {
	db := setupTestDB(t)
	tx := db

	reimbursementRepo := repositories.NewReimbursementRepository(tx)
	svc := services.NewService(tx, nil, nil, nil, nil, nil, reimbursementRepo)

	employeeId := uint(1)
	payrollId := uint(1)

	ctx := context.Background()
	req := models.EmployeeSubmitReimbursementRequest{
		EmployeeId:  employeeId,
		PayrollId:   payrollId,
		Amount:      10000,
		Description: "Test reimbursement",
	}

	_, err := svc.SubmitReimbursement(ctx, req)
	assert.NoError(t, err, "submit success")

}

package services

import (
	"context"
	"time"

	"github.com/payslip/models"
)

func (s *Service) SubmitReimbursement(ctx context.Context, req models.EmployeeSubmitReimbursementRequest) (interface{}, error) {
	var reimbursement models.Reimbursement
	reimbursement.EmployeeId = req.EmployeeId
	reimbursement.ReimbursementDate = time.Now().Format("2006-01-02")
	reimbursement.TotalAmount = req.Amount

	_, err := s.ReimbursementRepo.InsertReimbursement(ctx, &reimbursement)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

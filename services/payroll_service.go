package services

import (
	"context"
	"errors"
	"time"

	"github.com/payslip/models"
)

func (s *Service) AddPayroll(ctx context.Context, req models.AddPayrollRequest) (interface{}, error) {

	var payroll models.Payroll
	start_date, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}
	end_date, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, err
	}
	payroll.StartDate = start_date
	payroll.EndDate = end_date
	payroll.Description = req.Description

	//check payroll month already exist
	existing_payroll, err := s.PayrollRepo.FindPayrollByDate(ctx, &payroll)
	if err != nil {
		return nil, err
	}

	if existing_payroll != nil {
		return nil, errors.New("payroll already exist in this range of date")
	}

	err = s.PayrollRepo.CreatePayroll(ctx, &payroll)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

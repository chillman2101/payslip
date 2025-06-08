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
		return nil, errors.New("cannot parse")
	}
	end_date, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, errors.New("cannot parse")
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

func (s *Service) ListUnprocessedPayroll(ctx context.Context) (interface{}, error) {
	payrolls, err := s.PayrollRepo.ListPayrollUnprocessed(ctx)
	if err != nil {
		return nil, err
	}

	return payrolls, nil
}

func (s *Service) ProcessPayroll(ctx context.Context, payroll_id int) (interface{}, error) {
	err := s.PayrollRepo.ProcessPayroll(ctx, payroll_id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) GetSummaryPayrollEmployee(ctx context.Context, payroll_id, employee_id int) (interface{}, error) {
	payroll, err := s.PayrollRepo.GetSummaryPayrollByPayrollIdAndEmployeeId(ctx, payroll_id, employee_id)
	if err != nil {
		return nil, err
	}
	employee, err := s.AuthRepo.FindEmployeeById(ctx, uint(employee_id))
	if err != nil {
		return nil, err
	}

	// formatter response
	response := payroll.FormatterSummaryPayrollEmployee(employee)

	return response, nil
}

func (s *Service) GetSummaryPayrollAdmin(ctx context.Context, payroll_id int) (interface{}, error) {
	employees, err := s.AuthRepo.FindAllEmployee(ctx)
	if err != nil {
		return nil, err
	}

	var resp models.SummaryPayrollAdminResponse

	var totalTakeHomePay int64
	for _, employee := range employees {
		payroll, err := s.PayrollRepo.GetSummaryPayrollByPayrollIdAndEmployeeId(ctx, payroll_id, int(employee.ID))
		if err != nil {
			return nil, err
		}

		if len(payroll.Attendance) > 0 {
			// formatter response
			response := payroll.FormatterSummaryPayrollEmployee(&employee)
			totalTakeHomePay += response.TotalTakeHomePay
			resp.TakeHomePayEmployee = append(resp.TakeHomePayEmployee, models.TakeHomePayEmployee{
				EmployeeID:       int64(employee.ID),
				EmployeeName:     employee.Username,
				TotalTakeHomePay: response.TotalTakeHomePay,
			})
		}
	}
	resp.TotalTakeHomePayAllEmployee = totalTakeHomePay

	return resp, nil
}

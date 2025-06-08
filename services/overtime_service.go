package services

import (
	"context"
	"errors"
	"time"

	"github.com/payslip/models"
)

func (s *Service) SubmitOvertime(ctx context.Context, req models.EmployeeSubmitOvertimeRequest) (interface{}, error) {
	var overtime models.Overtime
	overtime.EmployeeId = req.EmployeeId
	overtime.OvertimeDate = time.Now().Format("2006-01-02")
	overtime.TotalHour = req.AmountTime
	overtime.PayrollId = req.PayrollId
	exist_overtime, err := s.OvertimeRepo.FindOvertimeByEmployeeIdAndDate(ctx, &overtime)
	if err != nil {
		return nil, err
	}

	if exist_overtime != nil {
		if exist_overtime.TotalHour+overtime.TotalHour > 3 {
			return nil, errors.New("Overtime cannot be requested up to 3 hours")
		}
		overtime.TotalHour += exist_overtime.TotalHour
		overtime.ID = exist_overtime.ID
		overtime.CreatedAt = exist_overtime.CreatedAt
	}

	_, err = s.OvertimeRepo.InsertOvertime(ctx, &overtime)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

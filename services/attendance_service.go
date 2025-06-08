package services

import (
	"context"
	"errors"
	"time"

	"github.com/payslip/models"
)

func (s *Service) CheckIn(ctx context.Context, req models.EmployeeAttendanceRequest) (interface{}, error) {
	var attendance models.Attendance
	attendance.EmployeeId = req.EmployeeId
	attendance.CheckInTime = time.Now()
	attendance.AttendanceDate = time.Now().Format("2006-01-02")
	exist_attendance, err := s.AttendanceRepo.FindAttendanceByEmployeeIdAndDate(ctx, &attendance)
	if err != nil {
		return nil, err
	}

	if exist_attendance != nil {
		return nil, errors.New("employee has already checked in")
	}

	_, err = s.AttendanceRepo.CheckInAttendance(ctx, &attendance)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) CheckOut(ctx context.Context, req models.EmployeeAttendanceRequest) (interface{}, error) {
	var attendance models.Attendance
	attendance.EmployeeId = req.EmployeeId
	attendance.AttendanceDate = time.Now().Format("2006-01-02")
	exist_attendance, err := s.AttendanceRepo.FindAttendanceByEmployeeIdAndDate(ctx, &attendance)
	if err != nil {
		return nil, err
	}

	if exist_attendance == nil {
		return nil, errors.New("employee has not checked in")
	}

	if !exist_attendance.CheckOutTime.IsZero() {
		return nil, errors.New("employee has already checked out")
	}

	exist_attendance.CheckOutTime = time.Now()
	_, err = s.AttendanceRepo.CheckOutAttendance(ctx, exist_attendance)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

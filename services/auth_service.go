package services

import (
	"context"
	"errors"

	"github.com/payslip/models"
	"github.com/payslip/utils"
)

func (s *Service) LoginAdmin(ctx context.Context, req models.UserRequest) (interface{}, error) {
	admin, err := s.AuthRepo.FindAdminByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	match := utils.CheckPasswordHash(req.Password, admin.Password)
	if !match {
		return nil, errors.New("invalid password")
	}

	return admin, nil
}

func (s *Service) LoginEmployee(ctx context.Context, req models.UserRequest) (interface{}, error) {
	employee, err := s.AuthRepo.FindEmployeeByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	match := utils.CheckPasswordHash(req.Password, employee.Password)
	if !match {
		return nil, errors.New("invalid password")
	}

	return employee, nil
}

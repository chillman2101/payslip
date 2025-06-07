package services

import (
	"context"

	"github.com/payslip/models"
)

func (s *Service) AuditLog(ctx context.Context, req models.AuditRequest) (interface{}, error) {
	s.Database.Table("audit_logs").Create(&req)
	return nil, nil
}

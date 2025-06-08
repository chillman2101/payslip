package services

import (
	"github.com/payslip/config/cache"
	"github.com/payslip/repositories"
	"gorm.io/gorm"
)

type Service struct {
	Database          *gorm.DB
	RedisCache        *cache.CacheHelper
	AuthRepo          repositories.AuthRepository
	PayrollRepo       repositories.PayrollRepository
	AttendanceRepo    repositories.AttendanceRepository
	OvertimeRepo      repositories.OvertimeRepository
	ReimbursementRepo repositories.ReimbursementRepository
}

func NewService(
	db *gorm.DB,
	redis *cache.CacheHelper,
	authRepo repositories.AuthRepository,
	payrollRepo repositories.PayrollRepository,
	attendanceRepo repositories.AttendanceRepository,
	overtimeRepo repositories.OvertimeRepository,
	reimbursementRepo repositories.ReimbursementRepository,
) *Service {
	return &Service{
		Database:          db,
		RedisCache:        redis,
		AuthRepo:          authRepo,
		PayrollRepo:       payrollRepo,
		AttendanceRepo:    attendanceRepo,
		OvertimeRepo:      overtimeRepo,
		ReimbursementRepo: reimbursementRepo,
	}
}

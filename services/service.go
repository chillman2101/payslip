package services

import (
	"github.com/payslip/config/cache"
	"github.com/payslip/repositories"
	"gorm.io/gorm"
)

type Service struct {
	Database    *gorm.DB
	RedisCache  *cache.CacheHelper
	AuthRepo    repositories.AuthRepository
	PayrollRepo repositories.PayrollRepository
}

func NewService(
	db *gorm.DB,
	redis *cache.CacheHelper,
	authRepo repositories.AuthRepository,
	payrollRepo repositories.PayrollRepository,
) *Service {
	return &Service{
		Database:    db,
		RedisCache:  redis,
		AuthRepo:    authRepo,
		PayrollRepo: payrollRepo,
	}
}

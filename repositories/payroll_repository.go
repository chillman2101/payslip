package repositories

import (
	"context"
	"errors"
	"log"
	"runtime/debug"

	"github.com/payslip/models"
	"gorm.io/gorm"
)

type PayrollRepository interface {
	FindPayrollByDate(ctx context.Context, payroll *models.Payroll) (*models.Payroll, error)
	CreatePayroll(ctx context.Context, payroll *models.Payroll) error
}

type payrollRepository struct {
	DB *gorm.DB
}

func NewPayrollRepository(db *gorm.DB) *payrollRepository {
	return &payrollRepository{DB: db}
}

func (pr *payrollRepository) FindPayrollByDate(ctx context.Context, payroll *models.Payroll) (*models.Payroll, error) {
	var exist_payroll models.Payroll
	// Find payroll by date
	query := pr.DB.Table("payrolls")
	query = query.Where("start_date <= ?", payroll.EndDate)
	query = query.Where("end_date >= ?", payroll.StartDate)
	err := query.First(&exist_payroll).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}

	return &exist_payroll, nil
}

func (pr *payrollRepository) CreatePayroll(ctx context.Context, payroll *models.Payroll) error {
	tx := pr.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Maaf gagal proses CreatePayroll: %v\n%s", r, debug.Stack())
		}
	}()

	err := tx.Create(payroll).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

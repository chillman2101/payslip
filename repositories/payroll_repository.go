package repositories

import (
	"context"
	"errors"
	"log"
	"runtime/debug"
	"time"

	"github.com/payslip/models"
	"gorm.io/gorm"
)

type PayrollRepository interface {
	FindPayrollByDate(ctx context.Context, payroll *models.Payroll) (*models.Payroll, error)
	CreatePayroll(ctx context.Context, payroll *models.Payroll) error
	ListPayrollUnprocessed(ctx context.Context) ([]models.Payroll, error)
	ProcessPayroll(ctx context.Context, payroll_id int) error
	GetSummaryPayrollByPayrollIdAndEmployeeId(ctx context.Context, payroll_id, employee_id int) (*models.Payroll, error)
	GetSummaryPayrollByPayrollId(ctx context.Context, payroll_id int) (*models.Payroll, error)
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

func (pr *payrollRepository) ListPayrollUnprocessed(ctx context.Context) ([]models.Payroll, error) {
	var payrolls []models.Payroll
	// Find payroll by date
	query := pr.DB.Table("payrolls")
	query = query.Where("already_proceed is false")
	err := query.Find(&payrolls).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}

	return payrolls, nil
}

func (pr *payrollRepository) ProcessPayroll(ctx context.Context, payroll_id int) error {
	tx := pr.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Maaf gagal proses CreatePayroll: %v\n%s", r, debug.Stack())
		}
	}()
	update := map[string]interface{}{
		"already_proceed": true,
		"updated_at":      time.Now(),
		"pay_date":        time.Now(),
	}
	err := tx.Model(&models.Payroll{}).Where("id = ?", payroll_id).Updates(update).Error
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

func (pr *payrollRepository) GetSummaryPayrollByPayrollIdAndEmployeeId(ctx context.Context, payroll_id, employee_id int) (*models.Payroll, error) {
	var payroll models.Payroll
	query := pr.DB.Table("payrolls")
	query = query.Where("id = ?", payroll_id)
	query = query.Preload("Attendance", "employee_id = ?", employee_id)
	query = query.Preload("Overtime", "employee_id = ?", employee_id)
	query = query.Preload("Reimbursement", "employee_id = ?", employee_id)

	err := query.First(&payroll).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}

	return &payroll, nil
}

func (pr *payrollRepository) GetSummaryPayrollByPayrollId(ctx context.Context, payroll_id int) (*models.Payroll, error) {
	var payroll models.Payroll
	query := pr.DB.Table("payrolls")
	query = query.Where("id = ?", payroll_id)
	query = query.Preload("Attendance")
	query = query.Preload("Overtime")
	query = query.Preload("Reimbursement")

	err := query.First(&payroll).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}

	return &payroll, nil
}

package repositories

import (
	"context"
	"errors"
	"log"
	"runtime/debug"

	"github.com/payslip/models"
	"gorm.io/gorm"
)

type OvertimeRepository interface {
	FindOvertimeByEmployeeIdAndDate(ctx context.Context, overtime *models.Overtime) (*models.Overtime, error)
	InsertOvertime(ctx context.Context, overtime *models.Overtime) (*models.Overtime, error)
}

type overtimeRepository struct {
	DB *gorm.DB
}

func NewOvertimeRepository(db *gorm.DB) *overtimeRepository {
	return &overtimeRepository{DB: db}
}

func (or *overtimeRepository) FindOvertimeByEmployeeIdAndDate(ctx context.Context, overtime *models.Overtime) (*models.Overtime, error) {
	// Find attendance by date
	var exist_overtime models.Overtime
	err := or.DB.Table("overtimes").Where("employee_id = ?", overtime.EmployeeId).Where("overtime_date::date = ?", overtime.OvertimeDate).First(&exist_overtime).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}

	return &exist_overtime, nil
}
func (or *overtimeRepository) InsertOvertime(ctx context.Context, overtime *models.Overtime) (*models.Overtime, error) {
	tx := or.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Maaf gagal proses CreatePayroll: %v\n%s", r, debug.Stack())
		}
	}()
	if overtime.ID == 0 {
		err := tx.Create(overtime).Error
		if err != nil {
			tx.Rollback()
			log.Println("DB error:", err)
			return nil, err
		}
	} else {
		err := tx.Save(overtime).Error
		if err != nil {
			tx.Rollback()
			log.Println("DB error:", err)
			return nil, err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		log.Println("DB error:", err)
		return nil, err
	}

	return overtime, nil
}

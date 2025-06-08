package repositories

import (
	"context"
	"errors"
	"log"
	"runtime/debug"

	"github.com/payslip/models"
	"gorm.io/gorm"
)

type ReimbursementRepository interface {
	FindReimbursementByEmployeeIdAndDate(ctx context.Context, reimbursement *models.Reimbursement) (*models.Reimbursement, error)
	InsertReimbursement(ctx context.Context, reimbursement *models.Reimbursement) (*models.Reimbursement, error)
}

type reimbursementRepository struct {
	DB *gorm.DB
}

func NewReimbursementRepository(db *gorm.DB) *reimbursementRepository {
	return &reimbursementRepository{DB: db}
}

func (or *reimbursementRepository) FindReimbursementByEmployeeIdAndDate(ctx context.Context, reimbursement *models.Reimbursement) (*models.Reimbursement, error) {
	// Find attendance by date
	var exist_reimbursement models.Reimbursement
	err := or.DB.Table("reimbursements").Where("employee_id = ?", reimbursement.EmployeeId).Where("reimbursement_date::date = ?", reimbursement.ReimbursementDate).First(&exist_reimbursement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}

	return &exist_reimbursement, nil
}
func (or *reimbursementRepository) InsertReimbursement(ctx context.Context, reimbursement *models.Reimbursement) (*models.Reimbursement, error) {
	tx := or.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Maaf gagal proses CreatePayroll: %v\n%s", r, debug.Stack())
		}
	}()

	err := tx.Create(reimbursement).Error
	if err != nil {
		tx.Rollback()
		log.Println("DB error:", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		log.Println("DB error:", err)
		return nil, err
	}

	return reimbursement, nil
}

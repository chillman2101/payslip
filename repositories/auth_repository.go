package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/payslip/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindEmployeeByUsername(ctx context.Context, username string) (*models.Employee, error)
	FindAdminByUsername(ctx context.Context, username string) (*models.Admin, error)
	FindEmployeeById(ctx context.Context, id uint) (*models.Employee, error)
	FindAllEmployee(ctx context.Context) ([]models.Employee, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{DB: db}
}

func (ar *authRepository) FindEmployeeByUsername(ctx context.Context, username string) (*models.Employee, error) {
	var employee models.Employee
	err := ar.DB.Where("username = ?", username).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found") // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}
	return &employee, nil
}

func (ar *authRepository) FindAdminByUsername(ctx context.Context, username string) (*models.Admin, error) {
	var admin models.Admin
	err := ar.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found") // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}
	return &admin, nil
}

func (ar *authRepository) FindEmployeeById(ctx context.Context, id uint) (*models.Employee, error) {
	var employee models.Employee
	err := ar.DB.Where("id = ?", id).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found") // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}
	return &employee, nil
}

func (ar *authRepository) FindAllEmployee(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	err := ar.DB.Find(&employees).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found") // Tidak ada konflik
		}
		log.Println("DB error:", err)
		return nil, err
	}
	return employees, nil
}

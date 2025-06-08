package database

import (
	"github.com/payslip/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Payroll{})
	db.AutoMigrate(&models.AuditLog{})
	db.AutoMigrate(&models.Attendance{})
	db.AutoMigrate(&models.Overtime{})
	db.AutoMigrate(&models.Reimbursement{})
}

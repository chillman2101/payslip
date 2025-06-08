package models

import (
	"gorm.io/gorm"
)

type Reimbursement struct {
	gorm.Model
	EmployeeId        uint   `json:"employee_id"`
	ReimbursementDate string `json:"reimbursement_date" gorm:"type:date;index"`
	TotalAmount       int64  `json:"total_hour"`
	Description       string `json:"description"`
	PayrollId         uint   `json:"payroll_id" gorm:"index"`
}

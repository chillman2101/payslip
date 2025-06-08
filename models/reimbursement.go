package models

import (
	"gorm.io/gorm"
)

type Reimbursement struct {
	gorm.Model
	EmployeeId        uint   `json:"employee_id"`
	ReimbursementDate string `json:"reimbursement_date"`
	TotalAmount       int64  `json:"total_hour"`
}

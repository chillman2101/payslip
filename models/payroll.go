package models

import (
	"time"

	"gorm.io/gorm"
)

type Payroll struct {
	gorm.Model
	Description    string          `json:"description"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	PayDate        time.Time       `json:"pay_date" gorm:"default:null"`
	AlreadyProceed bool            `json:"already_proceed" gorm:"default:false"`
	Attendance     []Attendance    `gorm:"foreignKey:PayrollId;references:ID"`
	Overtime       []Overtime      `gorm:"foreignKey:PayrollId;references:ID"`
	Reimbursement  []Reimbursement `gorm:"foreignKey:PayrollId;references:ID"`
}

func (Payroll) TableName() string {
	return "payrolls"
}

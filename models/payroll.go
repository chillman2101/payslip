package models

import (
	"time"

	"gorm.io/gorm"
)

type Payroll struct {
	gorm.Model
	Description    string    `json:"description"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	PayDate        time.Time `json:"pay_date" gorm:"default:null"`
	AlreadyProceed bool      `json:"already_proceed" gorm:"default:false"`
}

func (Payroll) TableName() string {
	return "payrolls"
}

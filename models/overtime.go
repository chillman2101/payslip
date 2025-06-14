package models

import (
	"gorm.io/gorm"
)

type Overtime struct {
	gorm.Model
	EmployeeId   uint   `json:"employee_id"`
	OvertimeDate string `json:"overtime_date" gorm:"type:date;index"`
	TotalHour    int64  `json:"total_hour"`
	PayrollId    uint   `json:"payroll_id" gorm:"index"`
}

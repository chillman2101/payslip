package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	EmployeeId     uint      `json:"employee_id"`
	AttendanceDate string    `json:"attendance_date" gorm:"type:date;index"`
	CheckInTime    time.Time `json:"check_in_time"`
	CheckOutTime   time.Time `json:"check_out_time"`
}

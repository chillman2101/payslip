package models

import "time"

type UserRequest struct {
	RequestId string `json:"request_id"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AddPayrollRequest struct {
	RequestId   string `json:"request_id"`
	Description string `json:"description" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"required"`
}

type AuditRequest struct {
	RequestId  string        `json:"request_id"`
	User       uint          `json:"user" binding:"required"`
	Endpoint   string        `json:"endpoint" binding:"required"`
	Method     string        `json:"method" binding:"required"`
	StatusCode int           `json:"status_code" binding:"required"`
	Duration   time.Duration `json:"duration" binding:"required"`
	ClientIp   string        `json:"client_ip" binding:"required"`
}

type EmployeeAttendanceRequest struct {
	RequestId  string `json:"request_id"`
	EmployeeId uint   `json:"employee_id"`
	PayrollId  uint   `json:"payroll_id"`
}

type EmployeeSubmitOvertimeRequest struct {
	RequestId   string `json:"request_id"`
	EmployeeId  uint   `json:"employee_id"`
	PayrollId   uint   `json:"payroll_id"`
	AmountTime  int64  `json:"amount_time" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type EmployeeSubmitReimbursementRequest struct {
	RequestId   string `json:"request_id"`
	EmployeeId  uint   `json:"employee_id"`
	PayrollId   uint   `json:"payroll_id"`
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
}

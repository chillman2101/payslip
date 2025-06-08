package models

import (
	"time"

	"github.com/payslip/utils"
)

type SummaryPayrollEmployeeResponse struct {
	EmployeeID       uint                  `json:"employee_id"`
	EmployeeName     string                `json:"employee_name"`
	Attendance       AttendanceResponse    `json:"attendance"`
	Overtime         OvertimeResponse      `json:"overtime"`
	Reimbursement    ReimbursementResponse `json:"reimbursement"`
	Salary           int64                 `json:"salary"`
	TotalTakeHomePay int64                 `json:"total_take_home_pay"`
}

type AttendanceResponse struct {
	Attendances     []ListAttendance `json:"attendances"`
	TotalDayPresent int64            `json:"total_day_present"`
	TotalWorkingDay int64            `json:"total_working_day"`
	Total           int64            `json:"total"`
}

type ListAttendance struct {
	AttendanceDate string `json:"attendance_date"`
	CheckInTime    string `json:"check_in_time"`
	CheckOutTime   string `json:"check_out_time"`
}

type OvertimeResponse struct {
	Overtimes         []ListOvertime `json:"overtimes"`
	TotalOvertimeHour int64          `json:"total_overtime_hour"`
	TotalWorkingHour  int64          `json:"total_working_hour"`
	Total             int64          `json:"total"`
}

type ListOvertime struct {
	OvertimeDate  string `json:"overtime_date"`
	OvertimeHours int64  `json:"overtime_hours"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}

type ReimbursementResponse struct {
	Reimbursements []ListReimbursement `json:"reimbursements"`
	Total          int64               `json:"total"`
}

type ListReimbursement struct {
	ReimbursementDate string `json:"reimbursement_date"`
	Amount            int64  `json:"amount"`
	ReimbursementNote string `json:"reimbursement_note"`
}

type SummaryPayrollAdminResponse struct {
	TakeHomePayEmployee         []TakeHomePayEmployee `json:"take_home_pay_employee"`
	TotalTakeHomePayAllEmployee int64                 `json:"total_take_home_pay_all_employee"`
}

type TakeHomePayEmployee struct {
	EmployeeID       int64  `json:"employee_id"`
	EmployeeName     string `json:"employee_name"`
	TotalTakeHomePay int64  `json:"total_take_home_pay"`
}

func (payroll *Payroll) FormatterSummaryPayrollEmployee(employee *Employee) *SummaryPayrollEmployeeResponse {
	var response SummaryPayrollEmployeeResponse
	response.EmployeeID = employee.ID
	response.EmployeeName = employee.Username
	response.Salary = employee.Salary

	totalDayPresent := 0
	totalWorkingDays := utils.CountWorkingDays(payroll.StartDate, payroll.EndDate)
	for _, attendance := range payroll.Attendance {
		if !attendance.CheckInTime.IsZero() && !attendance.CheckOutTime.IsZero() {
			totalDayPresent++
			response.Attendance.Attendances = append(response.Attendance.Attendances, ListAttendance{
				AttendanceDate: attendance.AttendanceDate,
				CheckInTime:    attendance.CheckInTime.Format("2006-01-02 15:04:05"),
				CheckOutTime:   attendance.CheckOutTime.Format("2006-01-02 15:04:05"),
			})
		}
	}

	// prorated salary calc
	proratedSalary := float64(employee.Salary) * (float64(totalDayPresent) / float64(totalWorkingDays))
	response.Attendance.Total = int64(proratedSalary)
	response.Attendance.TotalDayPresent = int64(totalDayPresent)
	response.Attendance.TotalWorkingDay = int64(totalWorkingDays)

	// Rate per hour calc
	totalWorkingHours := float64(totalWorkingDays * 8)
	ratePerHour := proratedSalary / float64(totalWorkingHours)

	// Overtime = 2x hourly rate
	totalOvertimeHour := 0
	for _, overtime := range payroll.Overtime {
		totalOvertimeHour += int(overtime.TotalHour)
		response.Overtime.Overtimes = append(response.Overtime.Overtimes, ListOvertime{
			OvertimeDate:  overtime.OvertimeDate,
			OvertimeHours: overtime.TotalHour,
			StartTime:     overtime.CreatedAt.Format("2006-01-02 15:04:05"),
			EndTime:       overtime.CreatedAt.Add(time.Hour * time.Duration(overtime.TotalHour)).Format("2006-01-02 15:04:05"),
		})
	}
	overtimeAmount := float64(totalOvertimeHour) * ratePerHour * 2
	response.Overtime.Total = int64(overtimeAmount)
	response.Overtime.TotalOvertimeHour = int64(totalOvertimeHour)
	response.Overtime.TotalWorkingHour = int64(totalWorkingHours)

	// Reimbursement Total
	var totalReimbursement float64
	for _, reimbursement := range payroll.Reimbursement {
		totalReimbursement += float64(reimbursement.TotalAmount)
		response.Reimbursement.Reimbursements = append(response.Reimbursement.Reimbursements, ListReimbursement{
			ReimbursementDate: reimbursement.ReimbursementDate,
			Amount:            reimbursement.TotalAmount,
			ReimbursementNote: reimbursement.Description,
		})
	}
	response.Reimbursement.Total = int64(totalReimbursement)

	// Total Take Home Pay
	totalTakeHomePay := proratedSalary + overtimeAmount + totalReimbursement
	response.TotalTakeHomePay = int64(totalTakeHomePay)

	return &response
}

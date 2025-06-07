package repositories

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/payslip/models"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	FindAttendanceByEmployeeIdAndDate(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error)
	CheckInAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error)
	CheckOutAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error)
}

type attendanceRepository struct {
	DB *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *attendanceRepository {
	return &attendanceRepository{DB: db}
}

func (ar *attendanceRepository) FindAttendanceByDate(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error) {
	// Find attendance by date
	var exist_attendance models.Attendance
	ar.DB.Table("attendances").Where("attendance_date::date = ?", attendance.AttendanceDate).First(&exist_attendance)

	return &exist_attendance, nil
}

func (ar *attendanceRepository) CheckInAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error) {
	tx := ar.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Maaf gagal proses CreatePayroll: %v\n%s", r, debug.Stack())
		}
	}()

	err := tx.Create(attendance).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return attendance, nil
}

func (ar *attendanceRepository) CheckOutAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error) {
	tx := ar.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Maaf gagal proses CreatePayroll: %v\n%s", r, debug.Stack())
		}
	}()

	err := tx.Save(attendance).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return attendance, nil
}

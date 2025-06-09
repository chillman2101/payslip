package integration_test

import (
	"context"
	"testing"

	"github.com/payslip/models"
	"github.com/payslip/repositories"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=test_payslip port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	return db
}

func TestEmployeeCheckIn_Integration(t *testing.T) {
	db := setupTestDB(t)
	tx := db

	attendanceRepo := repositories.NewAttendanceRepository(tx)
	svc := services.NewService(tx, nil, nil, nil, attendanceRepo, nil, nil)

	employeeId := uint(1)
	payrollId := uint(1)

	ctx := context.Background()
	req := models.EmployeeAttendanceRequest{
		EmployeeId: employeeId,
		PayrollId:  payrollId,
	}

	_, err := svc.CheckIn(ctx, req)
	assert.NoError(t, err, "check-in should succeed")

	_, err = svc.CheckIn(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "employee has already checked in", err.Error())

	_, err = svc.CheckOut(ctx, req)
	assert.NoError(t, err, "check-out should succeed")

	_, err = svc.CheckOut(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "employee has already checked out", err.Error())

}

func TestEmployeeCheckOut_WithoutCheckIn(t *testing.T) {
	db := setupTestDB(t)
	tx := db.Begin()
	defer tx.Rollback()

	attendanceRepo := repositories.NewAttendanceRepository(tx)
	svc := services.NewService(tx, nil, nil, nil, attendanceRepo, nil, nil)

	ctx := context.Background()
	req := models.EmployeeAttendanceRequest{
		EmployeeId: 9999,
		PayrollId:  1,
	}

	_, err := svc.CheckOut(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "employee has not checked in", err.Error())
}

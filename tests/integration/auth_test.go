package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/payslip/models"
	"github.com/payslip/repositories"
	"github.com/payslip/services"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLoginAdmin_Integration(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := &models.Admin{
		Username: "test_admin",
		Password: string(hashedPwd),
		Model:    gorm.Model{CreatedAt: time.Now()},
	}
	err := tx.Create(admin).Error
	assert.NoError(t, err)

	authRepo := repositories.NewAuthRepository(tx)
	svc := services.NewService(tx, nil, authRepo, nil, nil, nil, nil)

	ctx := context.Background()
	req := models.UserRequest{
		RequestId: uuid.NewString(),
		Username:  "test_admin",
		Password:  "admin123",
	}

	resp, err := svc.LoginAdmin(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)

	adminResp, ok := resp.(*models.Admin)
	assert.True(t, ok)
	assert.Equal(t, "test_admin", adminResp.Username)
}

func TestLoginAdmin_InvalidPassword_Integration(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=db_payslip port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	tx := db.Begin()
	defer tx.Rollback()

	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("correct_password"), bcrypt.DefaultCost)
	admin := &models.Admin{
		Username: "wrongpass_admin",
		Password: string(hashedPwd),
		Model:    gorm.Model{CreatedAt: time.Now()},
	}
	err = tx.Create(admin).Error
	assert.NoError(t, err)

	authRepo := repositories.NewAuthRepository(tx)
	svc := services.NewService(tx, nil, authRepo, nil, nil, nil, nil)

	ctx := context.Background()
	req := models.UserRequest{
		RequestId: uuid.NewString(),
		Username:  "wrongpass_admin",
		Password:  "wrong_password",
	}

	resp, err := svc.LoginAdmin(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "invalid password", err.Error())
}

func TestLoginAdmin_NotFound_Integration(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=db_payslip port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	tx := db.Begin()
	defer tx.Rollback()

	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("correct_password"), bcrypt.DefaultCost)
	admin := &models.Admin{
		Username: "wrong_username_admin",
		Password: string(hashedPwd),
		Model:    gorm.Model{CreatedAt: time.Now()},
	}
	err = tx.Create(admin).Error
	assert.NoError(t, err)

	authRepo := repositories.NewAuthRepository(tx)
	svc := services.NewService(tx, nil, authRepo, nil, nil, nil, nil)

	ctx := context.Background()
	req := models.UserRequest{
		RequestId: uuid.NewString(),
		Username:  "wrong_username_admin1",
		Password:  "wrong_password",
	}

	resp, err := svc.LoginAdmin(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "admin not found", err.Error())
}

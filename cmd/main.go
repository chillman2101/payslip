package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/payslip/config"
	"github.com/payslip/config/cache"
	"github.com/payslip/config/database"
	"github.com/payslip/config/redis"
	"github.com/payslip/middlewares"
	"github.com/payslip/repositories"
	"github.com/payslip/routes/admin"
	"github.com/payslip/routes/auth"
	"github.com/payslip/routes/employees/attendance"
	"github.com/payslip/routes/employees/overtime"
	"github.com/payslip/routes/employees/reimbursement"
	"github.com/payslip/services"
)

func main() {
	c, _ := config.LoadConfig()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// config database postgres
	database, err := database.NewDatabase(c)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("connected to database ")

	// config redis
	redisCache, err := redis.NewRedisCache(c)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	// cache helper
	cache_redis := cache.NewCacheHelper(redisCache)

	fmt.Println("connected redis cache & cache redis")

	//repositories
	authRepo := repositories.NewAuthRepository(database)
	payrollRepo := repositories.NewPayrollRepository(database)
	attendanceRepo := repositories.NewAttendanceRepository(database)
	overtimeRepo := repositories.NewOvertimeRepository(database)
	reimbursementRepo := repositories.NewReimbursementRepository(database)

	svc := services.NewService(database, cache_redis, authRepo, payrollRepo, attendanceRepo, overtimeRepo, reimbursementRepo)

	// setup routes
	r.Use(middlewares.RequestIDMiddleware())
	auth.RegisterRoutes(r, c, svc)
	admin.RegisterRoutes(r, c, svc)
	attendance.RegisterRoutes(r, c, svc)
	overtime.RegisterRoutes(r, c, svc)
	reimbursement.RegisterRoutes(r, c, svc)

	r.Run(fmt.Sprintf(":%s", c.ServerPort))
}

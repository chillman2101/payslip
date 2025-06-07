package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/payslip/middlewares"
	"github.com/payslip/models"
	"github.com/payslip/services"
)

func EmployeeLogin(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	var req models.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request", "request_id": requestID})
		return
	}

	result, err := svc.LoginEmployee(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	var employee models.Employee
	resp, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}
	json.Unmarshal(resp, &employee)

	token, err := middlewares.GenerateToken(employee.ID, "employee")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CheckIn(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")

	// Reject if weekend
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attendance cannot be submitted on weekends", "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully check in"})
}

func CheckOut(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")

	// Reject if weekend
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attendance cannot be submitted on weekends", "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully check out"})
}

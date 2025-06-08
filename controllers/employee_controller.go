package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/payslip/middlewares"
	"github.com/payslip/models"
	"github.com/payslip/services"
	"github.com/payslip/utils"
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
	user := c.GetUint("user_id")
	var req models.EmployeeAttendanceRequest
	req.RequestId = requestID
	req.EmployeeId = user
	now := time.Now().AddDate(0, 0, 1) // fake

	// Reject if weekend
	err := utils.ValidateOnlyWeekday(now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	_, err = svc.CheckIn(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully check in"})
}

func CheckOut(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	user := c.GetUint("user_id")
	var req models.EmployeeAttendanceRequest
	req.RequestId = requestID
	req.EmployeeId = user
	now := time.Now().AddDate(0, 0, 1) // fake

	// Reject if weekend
	err := utils.ValidateOnlyWeekday(now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	_, err = svc.CheckOut(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully check out"})
}

func SubmitOvertime(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	user := c.GetUint("user_id")

	var req models.EmployeeSubmitOvertimeRequest
	req.RequestId = requestID
	req.EmployeeId = user
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request", "request_id": requestID})
		return
	}
	now := time.Now() // fake
	fmt.Println("req", req)

	if now.Weekday() != time.Saturday && now.Weekday() != time.Sunday {
		// Reject if below 5 pm and weekday
		err := utils.ValidateOvertimeSubmission(now)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "request_id": requestID})
			return
		}
	}

	// reject if amount time > 3
	err := utils.ValidateOvertimeAmount(req.AmountTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	_, err = svc.SubmitOvertime(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully submit overtime"})
}

func SubmitReimbursement(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	user := c.GetUint("user_id")

	var req models.EmployeeSubmitReimbursementRequest
	req.RequestId = requestID
	req.EmployeeId = user
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request", "request_id": requestID})
		return
	}

	_, err := svc.SubmitReimbursement(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully submit reimbursement"})
}

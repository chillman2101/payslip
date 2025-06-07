package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/payslip/middlewares"
	"github.com/payslip/models"
	"github.com/payslip/services"
)

func AdminLogin(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	var req models.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request", "request_id": requestID})
		return
	}

	result, err := svc.LoginAdmin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	var admin models.Admin
	resp, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}
	json.Unmarshal(resp, &admin)

	token, err := middlewares.GenerateToken(admin.ID, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func AddPayroll(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	var req models.AddPayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request :" + err.Error(), "request_id": requestID})
		return
	}

	_, err := svc.AddPayroll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully Create Payroll"})
}

package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/payslip/models"
	"github.com/payslip/services"
)

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

func ListUnprocessedPayroll(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")

	result, err := svc.ListUnprocessedPayroll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	var payrolls []models.Payroll
	resp, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}
	json.Unmarshal(resp, &payrolls)

	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully Get Unprocessed Payroll", "data": &payrolls})
}

func ProcessPayroll(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	var payroll_id int
	id, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithError(http.StatusBadRequest, errors.New("id is required"))
		return
	}
	payroll_id, _ = strconv.Atoi(id)

	_, err := svc.ProcessPayroll(c.Request.Context(), payroll_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully Processed Payroll"})
}

func SummaryPayroll(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	user := c.GetUint("user_id")
	var payroll_id int
	id, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithError(http.StatusBadRequest, errors.New("id is required"))
		return
	}
	payroll_id, _ = strconv.Atoi(id)

	response, err := svc.GetSummaryPayrollEmployee(c.Request.Context(), payroll_id, int(user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully get summary", "data": &response})
}

func SummaryPayrollAdmin(c *gin.Context, svc *services.Service) {
	requestID := c.GetString("request_id")
	var payroll_id int
	id, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithError(http.StatusBadRequest, errors.New("id is required"))
		return
	}
	payroll_id, _ = strconv.Atoi(id)

	response, err := svc.GetSummaryPayrollAdmin(c.Request.Context(), payroll_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "request_id": requestID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully get summary", "data": &response})
}

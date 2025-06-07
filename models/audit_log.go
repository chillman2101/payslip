package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	gorm.Model
	RequestId  string        `json:"request_id"`
	User       uint          `json:"user"`
	Endpoint   string        `json:"endpoint"`
	Method     string        `json:"method"`
	StatusCode int           `json:"status_code"`
	Duration   time.Duration `json:"duration"`
	ClientIp   string        `json:"client_ip"`
}

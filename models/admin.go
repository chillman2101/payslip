package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex:idx_username"`
	Password string `json:"password"` // exclude from SQL generation
}

func (Admin) TableName() string {
	return "admins"
}

var Admins = []UsernameDummy{
	{Username: "admin"},
	{Username: "admin123"},
}

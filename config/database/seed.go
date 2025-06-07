package database

import (
	"fmt"

	"github.com/payslip/models"
	"github.com/payslip/utils"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	err := SeederGenerateEmployee(db)
	if err != nil {
		fmt.Println("employee seeder error: ", err)
	}

	err = SeederGenerateAdmin(db)
	if err != nil {
		fmt.Println("employee seeder error: ", err)
	}
}

func SeederGenerateEmployee(db *gorm.DB) error {
	users := models.Users
	for _, user := range users {
		employee := models.Employee{
			Username: user.Username,
			Password: utils.HashPassword(user.Username),
			Salary:   int64(utils.RandomSalary()),
		}
		err := db.Create(&employee).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func SeederGenerateAdmin(db *gorm.DB) error {
	admins := models.Admins
	for _, admin := range admins {
		admin_user := models.Admin{
			Username: admin.Username,
			Password: utils.HashPassword(admin.Username),
		}
		err := db.Create(&admin_user).Error
		if err != nil {
			return err
		}
	}
	return nil
}

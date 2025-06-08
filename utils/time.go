package utils

import (
	"errors"
	"time"
)

func ValidateOvertimeSubmission(submissionTime time.Time) error {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	workEnd := time.Date(
		submissionTime.Year(),
		submissionTime.Month(),
		submissionTime.Day(),
		17, 0, 0, 0, loc,
	)

	if submissionTime.Before(workEnd) {
		return errors.New("Overtime can only be requested after 5pm")
	}
	return nil
}

func ValidateOnlyWeekday(attendanceTime time.Time) error {
	if attendanceTime.Weekday() == time.Saturday || attendanceTime.Weekday() == time.Sunday {
		return errors.New("Attendance cannot be submitted on weekends")
	}
	return nil
}

func ValidateOvertimeAmount(amountTime int64) error {
	if amountTime > 3 {
		return errors.New("Overtime cannot be requested up to 3 hours")
	}
	return nil
}

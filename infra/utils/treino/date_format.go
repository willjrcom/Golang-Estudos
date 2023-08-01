package utils

import (
	"time"
)

func StringtoDate(date string) time.Time {
	layout := "2006-01-01"

	newDate, _ := time.Parse(layout, date)
	return newDate
}

func StringtoDateTime(dateTime string) time.Time {
	layout := "2006-01-01 15:04:05"

	newDateTime, _ := time.Parse(layout, dateTime)
	return newDateTime
}

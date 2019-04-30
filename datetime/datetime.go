package datetime

import (
	"time"
)

const (
	DateTimeFormatter = "20060102150405"
	DateFormatter     = "20060102"
	TimeFormatter     = "150405"
	Year              = "2006"
	Month             = "01"
	Day               = "02"
	Hour              = "15"
	Minute            = "04"
	Second            = "05"
)

func CurrentDateTimeStr() string {
	return time.Now().Format(DateTimeFormatter)
}

func CurrentDateStr() string {
	return time.Now().Format(DateFormatter)
}

func CurrentTimeStr() string {
	return time.Now().Format(TimeFormatter)
}

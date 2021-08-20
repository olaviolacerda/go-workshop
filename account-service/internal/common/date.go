package common

import "time"

func DateToStr(value time.Time) string {
	format := "01/02/2006 15:04:05"
	return value.Format(format)
}

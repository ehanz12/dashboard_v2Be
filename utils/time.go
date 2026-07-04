package utils

import "time"

var Jakarta, _ = time.LoadLocation("Asia/Jakarta")

func Now() time.Time {
	return time.Now().In(Jakarta)
}

func TodayString() string {
	return Now().Format("2006-01-02")
}

func TimeString() string {
	return Now().Format("15:04")
}

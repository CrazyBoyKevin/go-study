package models

import (
	"time"
)

// UnixToTime 时间戳转日期
func UnixToTime(timestamp int64) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// DateToUnix 日期转时间戳
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	location, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	} else {
		return location.Unix()
	}
}

// GetUnix 获取当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetTime 获取当前时间
func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetDate 获取当前日期
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

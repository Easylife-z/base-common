package utils

import (
	"strconv"
	"time"
)

// 获取当月第一天
func GetCurrentMonthFirstDay() time.Time {
	now := time.Now()
	firstdayTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	return firstdayTime
}

// 获取当前月份，格式：202403，一般用于月统计数据，数据库存该格式字段
func GetCurrentMonth() int64 {
	timeStr := time.Now().Format("200601")

	if month, err := strconv.Atoi(timeStr); err == nil {
		return int64(month)
	}
	return 0
}

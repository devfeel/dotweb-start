package _time

import (
	"errors"
	"time"
)

//ParseTime 使用常用的时间格式尝试将字符串转换为time对象
func ParseTime(timeStr string) (time.Time, error) {
	timeFormatList := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-1-2 15:04:05",
		"2006-1-2 15:04",
		"2006-1-2 15:04 MST",
		"2006-1-2 15:04:05 MST",
		"2006-01-02 15:04 MST",
		"2006-01-02 15:04:05 MST",
		"2006-01-02T15:04:05Z",
		"2006-1-2T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-1-2T15:04:05",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04",
		"2006/1/2 15:04:05",
		"2006/1/2 15:04",
		"2006/1/2 15:04 MST",
		"2006/1/2 15:04:05 MST",
		"2006/01/02 15:04 MST",
		"2006/01/02 15:04:05 MST",
		"2006/01/02T15:04:05Z",
		"2006/1/2T15:04:05Z",
		"2006-1-2",
		"2006-01-02"}
	for _, format := range timeFormatList {
		time, err := time.ParseInLocation(format, timeStr, time.Local)
		if err == nil {
			return time, nil
		}
	}
	return time.Time{}, errors.New("未找到匹配的时间格式")

}

// ToDay 获取当前日期
//一个对象，设置为当天日期，其时间组成部分设置为 00:00:00。
func ToDay() time.Time {
	now := time.Now()
	time, _ := time.ParseInLocation("2006-1-2 MST", now.Format("2006-1-2 MST"), time.Local)
	return time
}

// GetMondayAndFriday 获取指定日期所在周的周一和周五
func GetMondayAndFriday(date time.Time) (monday time.Time, friday time.Time) {
	weekDay := int(date.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	monday = date.AddDate(0, 0, -weekDay+1)
	friday = date.AddDate(0, 0, 5-weekDay)
	return monday, friday
}

// GetTimestamp 获取时间戳 毫秒级
func GetTimestamp(date time.Time) (timestamp int64) {
	timestamp = date.UnixNano() / 1000000
	return timestamp
}

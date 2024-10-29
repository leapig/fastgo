package utils

import (
	"strings"
	"time"
)

// 时间基准点
var timePoint = "2006-01-02 15:04:05"

// TimeToString time.Time转string
func TimeToString(t time.Time) string {
	if &t == nil {
		return ""
	} else {
		return t.Format(timePoint)
	}
}

var dateTimePoint = "2006-01-02"

func DateTimeToString(t time.Time) string {
	if &t == nil {
		return ""
	} else {
		return t.Format(dateTimePoint)
	}
}

func StringToDateTime(str string) (time.Time, error) {
	return time.ParseInLocation(dateTimePoint, str, time.Local)
}

// StrToTime string转time.Time
func StrToTime(str string) (time.Time, error) {
	t, _ := time.ParseInLocation(timePoint, str, time.Local)
	return t, nil
}

func StartTime(startTime string) (time.Time, error) {
	//t, _ := time.ParseInLocation(timePoint, startTime, time.Local)
	//if !t.IsZero() {
	//pattern := `^d{4}-d{2}-d{2}$`
	//result, _ := regexp.MatchString(pattern, startTime)
	//if result {
	tt, _ := time.ParseInLocation(timePoint, startTime+" 00:00:00", time.Local)
	return tt, nil
	//} else {
	//	t, _ := time.ParseInLocation(timePoint, startTime, time.Local)
	//	return t, nil
	//}
}

func EndTime(endTime string) (time.Time, error) {
	//pattern := `^d{4}-d{2}-d{2}$`
	//result, _ := regexp.MatchString(pattern, endTime)
	//if result {
	tt, _ := time.ParseInLocation(timePoint, endTime+" 23:59:59", time.Local)
	return tt, nil
	//} else {
	//	t, _ := time.ParseInLocation(timePoint, endTime, time.Local)
	//	return t, nil
	//}

}

func TodayStartTime(time time.Time) (time.Time, error) {
	str := Now()
	return StrToTime(strings.Split(str, " ")[0] + " 00:00:00")
}

func TodayEndTime(time time.Time) (time.Time, error) {
	str := Now()
	return StrToTime(strings.Split(str, " ")[0] + " 23:59:59")
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowYYMMDD() string {
	return time.Now().Format("2006-01-02")
}

func CheckOnlineDevice(te time.Time, str string) bool {
	m, _ := time.ParseDuration("-" + str + "m")
	result := time.Now().Add(m)
	return result.Before(te)
}

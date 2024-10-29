package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// RandomStr 获取指定位数的随机数
func RandomStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomStrByTime 获取指定位数的随机数(用当前时间作为随机因子)
func RandomStrByTime(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// SHA1 SHA1加密
func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}

// IntToString int转string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString Int64转string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// StringToInt string转int
func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// StringToInt64 string装int64
func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func Float64ToString(f float64) string {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	return str
}

// ConvertToStringSlice 将int64类型的切片转换为string类型的切片
func ConvertToStringSlice(arr []int64) []string {
	var strArr []string
	for _, num := range arr {
		strNum := strconv.FormatInt(num, 10) // 将int64类型的数值转换为字符串
		strArr = append(strArr, strNum)      // 添加到新的切片中
	}
	return strArr
}

// RemoveDuplicates 去除重复字符串
func RemoveDuplicates(a []string) []string {
	m := make(map[string]bool)
	result := make([]string, 0)
	for _, str := range a {
		if !m[str] {
			m[str] = true
			result = append(result, str)
		}
	}
	return result
}

func ParseIdCard(idCard string) (time.Time, int32) {
	year := idCard[6:10]
	mouth := idCard[10:12]
	day := idCard[12:14]
	gender := idCard[16:17]
	birthday, _ := StringToDateTime(year + "-" + mouth + "-" + day)
	if StringToInt64(gender)%2 == 0 {
		return birthday, 2
	} else {
		return birthday, 1
	}
}

package util

import (
	"math"
	"reflect"
	"strconv"
	"time"
)

const (
	Data_Format_Template = "2006-01-02 15:04:05"
	One_Day_Second       = 24 * 3600
)

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

//func Typeof(v interface{}) string {
//	return fmt.Sprintf("%T", v)
//}

func GetMsgName(msg interface{}) string {
	msgType := reflect.TypeOf(msg)
	msgName := msgType.Elem().Name()
	return msgName
}

func B2s(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}
func IntToString(c int) string {
	return strconv.Itoa(c)
}
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
func Int64ToString(str int64) string {
	return strconv.FormatInt(str, 10)
}
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
func StringToInt32(str string) (int32, error) {
	v, err := strconv.ParseInt(str, 10, 32)
	return int32(v), err
}
func StringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

//---------------------------------------
//将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

//将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

//精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

//精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}

//---------------------------------------
func GetCurrentTimeByMS() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTodayZeorTime() int64 {
	t := time.Now()
	year, month, day := t.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return today.Unix()
}

func GetTodayOffHour(offHour int) int64 {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, offHour, 0, 0, 0, t.Location()).Unix()
}

func DateStr2Unix(dateStr string) int64 {
	loc, _ := time.LoadLocation("Local") //取本地时区
	date, err := time.ParseInLocation(Data_Format_Template, dateStr, loc)
	if err != nil {
		return 0
	}
	return date.Unix()
}

func Unix2DateStr(unix int64) string {
	strTime := time.Unix(unix, 0).Format(Data_Format_Template)
	return strTime
}

func Int642String(n int64) string {
	return strconv.FormatInt(n, 10)
}

func String2Int64(n string) int64 {
	int64, _ := strconv.ParseInt(n, 10, 64)
	return int64
}

func String2Float64(n string) float64 {
	v, _ := strconv.ParseFloat(n, 64)
	return v
}

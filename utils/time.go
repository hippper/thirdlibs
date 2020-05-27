package utils

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT_WITH_T   = "2006-01-02T15:04:05"
	TIME_FORMAT1         = "02-01-2006 15:04:00" // "dd-MM-yyyy HH:mm:ss" 秒数为 00
	TIME_FORMAT2         = "02-01-2006 15:04:05" // "dd-MM-yyyy HH:mm:ss"
	TIME_FORMAT3         = "01/02/2006 15:04:05" // MM/DD/YYYY HH:MM:SS
	TIME_FORMAT4         = "2006-01-02 15:04:00" // 秒数为 00
	TIME_FORMAT5         = "2006/01/02 15:04:05" // YYYY/MM/DD HH:mm:ss
	TIME_FORMAT_COMPACT1 = "060102150405"
	DATE_FORMAT1         = "2006/01/02"
	DATE_FORMAT_COMPACT1 = "20060102 15:04:05"
	MONTH_FORMAT1        = "2006-1"

	TIME_FORMAT_WITH_MS         = "2006-01-02 15:04:05.000"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	TIME_FORMAT_WO_SEC_COMPACT  = "200601021504"
	TIME_FORMAT_COMPACT         = "20060102150405"
	TIME_FORMAT_WITH_MS_COMPACT = "20060102150405000"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT_COMPACT         = "20060102"
	MONTH_FORMAT                = "2006-01"
)

const (
	TIME_LOC_ASIA_SHANGHAI    = "Asia/Shanghai"    //+0800
	TIME_LOC_ASIA_TAIPEI      = "Asia/Taipei"      //+0800
	TIME_LOC_AMERICA_NEW_YORK = "America/New_York" //-0400(夏) -0500(冬）
	TIME_LOC_UTC              = "UTC"
)

func TimeStr(t time.Time) string {
	return t.Format(TIME_FORMAT)
}

func DateStr(t time.Time) string {
	return t.Format(DATE_FORMAT)
}

func NowStr() string {
	timenow := time.Now().Format(TIME_FORMAT)
	return timenow
}

func NowWithMS() string {
	timenow := time.Now().Format(TIME_FORMAT_WITH_MS_COMPACT)
	return timenow
}

func NowMS() string {
	timenow := time.Now().Format(TIME_FORMAT_WITH_MS)
	return timenow
}

func NowDate() string {
	timenow := time.Now().Format(DATE_FORMAT_COMPACT)
	return timenow
}

func NowDateStr() string {
	timenow := time.Now().Format(DATE_FORMAT)
	return timenow
}

// Millisecond timestrap
func GetTimeOfMs() int64 {
	return time.Now().UnixNano() / 1000000
}

// Nanosecond timestrap
func GetTimeOfNs() int64 {
	return time.Now().UnixNano()
}

// Second timestrap
func GetTimeOfS() int64 {
	return time.Now().Unix()
}

// string time to time.Time in location
func GetTimeFromFormat(layout string, timeStr string, location string) (time.Time, error) {
	var tTime time.Time

	loc, err := time.LoadLocation(location)
	if err != nil {
		return tTime, err
	}

	tTime, err = time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		return tTime, err
	}

	return tTime, nil
}

// time.Time to string time in location
func GetTimeStrFromFormat(layout string, timeClass time.Time, location string) string {
	local, err := time.LoadLocation(location)
	if nil != err {
		return timeClass.Format(layout)
	}
	return timeClass.In(local).Format(layout)
}

// change time zone
func GetTimeByChangeTimeZone(layout string, timeStr string, srcTimeZone string, destTimeZone string) (time.Time, error) {
	var timestamp time.Time

	localTimeStr, err := time.LoadLocation(srcTimeZone)
	if err != nil {
		return timestamp, err
	}
	timestamp, err = time.ParseInLocation(layout, timeStr, localTimeStr)
	if err != nil {
		return timestamp, err
	}

	local, err := time.LoadLocation(destTimeZone)
	if err != nil {
		return timestamp, err
	}

	return timestamp.In(local), nil
}

// get the first date and the last date of a month
func GetOneMonthRangeDateStr(d time.Time) (string, string, error) {
	fdatestr := fmt.Sprintf("%04d%02d%02d000000", d.Year(), d.Month(), 01)

	fdate, err := time.Parse(TIME_FORMAT_COMPACT, fdatestr)
	if err != nil {
		return "", "", err
	}

	ldate := fdate.AddDate(0, 1, 0).AddDate(0, 0, -1)

	return DateStr(fdate), DateStr(ldate), nil
}

// get the first date and the last date of a month
func GetOneMonthRangeDate(d time.Time) (time.Time, time.Time, error) {
	fdatestr := fmt.Sprintf("%04d%02d%02d000000", d.Year(), d.Month(), 01)

	fdate, err := time.Parse(TIME_FORMAT_COMPACT, fdatestr)
	if err != nil {
		return fdate, fdate, err
	}

	ldate := fdate.AddDate(0, 1, 0).AddDate(0, 0, -1)

	return fdate, ldate, nil
}

func IsValidDateTime(dt string) bool {
	_, err := time.Parse(TIME_FORMAT, dt)
	if err != nil {
		return false
	}
	return true
}

func IsValidDate(d string) bool {
	_, err := time.Parse(TIME_FORMAT, d+" 00:00:00")
	if err != nil {
		return false
	}
	return true
}

func IsValidTime(t string) bool {
	dt := fmt.Sprintf("2012-12-12 %s", t)

	return IsValidDateTime(dt)
}

func IsNullTime(t time.Time) bool {
	year := t.Year()
	if year == 1 {
		return true
	} else {
		return false
	}
}

func GetTimeFromString(timeStr string) (time.Time, error) {

	tTime, err := time.ParseInLocation(TIME_FORMAT, timeStr, time.Local)
	if err != nil {
		return tTime, err
	}

	return tTime, nil
}

func Time2Date(t time.Time) time.Time {
	timedate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return timedate
}

func MakeEternalTime() time.Time {
	return time.Date(9999, 12, 31, 0, 0, 0, 0, time.Local)
}

func CalcTimeSecond(stime string) (*time.Time, int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return nil, 0, err
	}

	ltime, err := time.ParseInLocation(TIME_FORMAT_COMPACT, stime, loc)
	if err != nil {
		return nil, 0, err
	}

	return &ltime, ltime.Unix(), nil
}

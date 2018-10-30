package utils

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT_WITH_MS         = "2006-01-02 15:04:05.000"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	TIME_FORMAT_COMPACT         = "20060102150405"
	TIME_FORMAT_WITH_MS_COMPACT = "20060102150405.000"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT_COMPACT         = "20060102"
	MONTH_FORMAT                = "2006-01"
)

const (
	TIME_LOC_ASIA_SHANGHAI    = "Asia/Shanghai"    //+0800
	TIME_LOC_ASIA_TAIPEI      = "Asia/Taipei"      //+0800
	TIME_LOC_AMERICA_NEW_YORK = "America/New_York" //-0400(夏) -0500(冬）
)

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

// time string to timestrap
func GetTimeMsFromFormat(layout string, timeStr string, location string) (int64, error) {

	var timestamp time.Time
	local, err := time.LoadLocation(location)
	if nil != err {
		fmt.Printf("LoadLocation Failed then set local time, locationArg :%s, err :%v", location, err)
		timestamp, err = time.Parse(layout, timeStr)
	} else {
		timestamp, err = time.ParseInLocation(layout, timeStr, local)
	}

	return timestamp.UnixNano() / 1000000, err
}

// timestrap to fromat time string
func Timestamp2str(timestamp int64, layout string) string {
	strTime := ""
	if timestamp > 0 {
		strTime = time.Unix(timestamp, 0).Format(layout)
	}
	return strTime
}

// string time to time.Time in location
func GetTimeFromFormat(layout string, timeStr string, location string) (time.Time, error) {
	var timestamp time.Time
	local, err := time.LoadLocation(location)
	if nil != err {
		fmt.Printf("LoadLocation Failed then set local time, locationArg :%s, err :%v", location, err)
		timestamp, err = time.Parse(layout, timeStr)
	} else {
		timestamp, err = time.ParseInLocation(layout, timeStr, local)
	}

	return timestamp, err
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

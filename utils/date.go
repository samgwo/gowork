package utils

import (
	//"fmt"
	"time"
)

/*
func GetDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}
func GetDateMH(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01-02 03:04")
}*/

func GetDateFormat(timestamp int64, format string) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func GetDate(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

func GetDateMH(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}

func GetTimeParse(times string) int64 {
	if "" == times {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02 15:04", times, loc)
	return parse.Unix()
}

func GetDateParse(dates string) int64 {
	if "" == dates {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02", dates, loc)
	return parse.Unix()
}

func CheckWorkTime(wk int, st int64, et int64) bool {
	r := false
	var d int
	d = 0
	if et >= st {
		d = (int((et-st)/(3600*24)) + 1) * 8
	}
	if wk <= d && wk > 0 {
		r = true
	}
	return r
}

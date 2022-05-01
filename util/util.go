package util

import (
	"fmt"
	"strconv"
	"time"
)

func Assert(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func formatNumber(n int) string {
	s := strconv.Itoa(n)
	if len(s) == 1 {
		return "0" + s
	}
	return s
}

func formatTimeToTimeStamp(timeStamp int64) string {
	date := time.Unix(timeStamp, 0)
	timeString := fmt.Sprintf("%s%s%s%s%s%s",
		formatNumber(int(date.Month())),
		formatNumber(date.Day()),
		formatNumber(date.Year()),
		formatNumber(date.Hour()),
		formatNumber(date.Minute()),
		formatNumber(date.Second()))
	return timeString
}

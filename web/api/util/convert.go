package util

import (
	"strconv"
	"time"
)

func ConvertStringToUint(s string) uint {
	uint64, _ := strconv.ParseUint(s, 10, 32)
	return uint(uint64)
}

func ConvertStringToTime(str string) time.Time {
	t, _ := time.Parse("2022-01-01 00:00:00", str)
	return t
}

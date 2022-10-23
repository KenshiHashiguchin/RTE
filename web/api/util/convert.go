package util

import "strconv"

func ConvertStringToUint(s string) uint {
	uint64, _ := strconv.ParseUint(s, 10, 32)
	return uint(uint64)
}

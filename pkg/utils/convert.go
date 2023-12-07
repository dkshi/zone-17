package utils

import "strconv"

func UInt64ToString(x uint64) string {
	return strconv.FormatUint(x, 10)
}
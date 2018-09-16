package util

import (
	"strconv"
)

// CheckError panics if err is not nil.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// ParseInt returns a base 10 int32
func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	CheckError(err)
	
	return int(i)
}
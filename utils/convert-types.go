package utils

import (
	"strconv"
	"time"
)

func fromStringToBool(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolValue
}

// ConvertToBoolArray converts a string array to a bool array
func ConvertToBoolArray(value []string) []bool {
	length := len(value)
	var boolArray []bool
	for i := 0; i < length; i++ {
		boolArray[i] = fromStringToBool(value[i])
	}
	return boolArray
}

// ConvertToInt converts a string to an int
func ConvertToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

// ConvertToUInt converts a string to a uint
func ConvertToUInt(value string) uint {
	uintValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0
	}
	return uint(uintValue)
}

func ConvertToTime(value string) time.Time {
	intValue := ConvertToInt(value)
	return time.Unix(int64(intValue), 0)
}

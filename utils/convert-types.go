package utils

import (
	utilsDiagnostics "backend/diagnostics/utils"
	"github.com/jinzhu/now"
	"strconv"
	"strings"
	"time"
)

// ConvertToBool converts a string to a bool
func ConvertToBool(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		utilsDiagnostics.ConvertToBoolErr(err, value)
		return false
	}
	return boolValue
}

// ConvertToBoolArray converts a string array to a bool array
func ConvertToBoolArray(value []string) []bool {
	length := len(value)
	boolArray := make([]bool, length)
	for i := 0; i < length; i++ {
		boolArray[i] = ConvertToBool(value[i])
	}
	return boolArray
}

// ConvertToInt converts a string to an int
func ConvertToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		utilsDiagnostics.ConvertToIntErr(err, value)
		return 0
	}
	return intValue
}

// ConvertToUInt converts a string to a uint
func ConvertToUInt(value string) uint {
	uintValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		utilsDiagnostics.ConvertToUIntErr(err, value)
		return 0
	}
	return uint(uintValue)
}

// ConvertToTime converts an int timestamp to time object
func ConvertToTime(value string) time.Time {
	intValue := ConvertToInt(value)
	intValue /= 1000
	return time.Unix(int64(intValue), 0)
}

// ConvertToStartOfDay converts the string time to a time object (at the beginning of the day)
func ConvertToStartOfDay(value string) time.Time {
	date := ConvertToTime(value)
	return now.With(date).BeginningOfDay()
}

// ConvertToUIntArray converts the string array to uint array
func ConvertToUIntArray(value []string) []uint {
	length := len(value)
	uintArray := make([]uint, length)
	for i := 0; i < length; i++ {
		uintArray[i] = ConvertToUInt(value[i])
	}
	return uintArray
}

// ConvertToFormArray returns an array from the form recieved variable
func ConvertToFormArray(value string) []string {
	if len(value) == 0 {
		return make([]string, 0)
	}
	return strings.Split(value, ",")
}
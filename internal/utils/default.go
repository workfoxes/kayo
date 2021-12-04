package utils

import (
	"strconv"

	"github.com/workfoxes/calypso/pkg/log"
)

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}
	log.Panic("ParseFloat: ", s, " -> ", err.Error())
	return 0
}

// ParseFloatIgnoreError is like ParseFloat but ignores errors.
func ParseFloatIgnoreError(s string, defaultValue float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}
	log.Error("ParseFloat: ", s, " -> ", err.Error())
	return defaultValue
}

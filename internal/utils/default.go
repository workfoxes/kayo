package utils

import (
	"strconv"

	"github.com/workfoxes/kayo/pkg/log"
)

func ParseFloat(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	} else {
		log.Error("ParseFloat: %s -> %s", s, err.Error())
		return 0
	}
}

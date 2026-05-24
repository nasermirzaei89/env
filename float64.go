package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetFloat64 extracts float64 value from env. If not set, returns default value.
func GetFloat64(key string, def float64) float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

// MustGetFloat64 extracts float64 value from env. If not set, it panics.
func MustGetFloat64(key string) float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

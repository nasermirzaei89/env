package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetFloat64 extracts int value from env. if not set, returns default value.
func GetFloat64(key string, def float64) float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		return def
	}

	return v
}

// MustGetFloat64 extracts int value from env. if not set, it panics.
func MustGetFloat64(key string) float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return v
}

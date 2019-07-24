package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetFloat extracts int value from env
// if not set, returns default value
func GetFloat(key string, def float64) float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}

	return v
}

// MustGetFloat extracts int value from env
// if not set, it panics
func MustGetFloat(key string) float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return v
}

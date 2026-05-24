package env

import (
	"fmt"
	"os"
	"time"
)

// GetDuration extracts duration value from env. If not set, returns default value.
func GetDuration(key string, def time.Duration) time.Duration {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	res, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has invalid duration value: %v", key, err))
	}

	return res
}

// MustGetDuration extracts duration value from env. If not set, it panics.
func MustGetDuration(key string) time.Duration {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	res, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has invalid duration value: %v", key, err))
	}

	return res
}

package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt extracts int value from env. If not set, returns default value.
func GetInt(key string, def int) int {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

// MustGetInt extracts int value from env. If not set, it panics.
func MustGetInt(key string) int {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

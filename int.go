package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt extracts int value from env. if not set, returns default value.
func GetInt(key string, def int) int {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}

	return v
}

// MustGetInt extracts int value from env. if not set, it panics.
func MustGetInt(key string) int {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return v
}

package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt64 extracts int64 value from env. If not set, returns default value.
func GetInt64(key string, def int64) int64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

// MustGetInt64 extracts int64 value from env. If not set, it panics.
func MustGetInt64(key string) int64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

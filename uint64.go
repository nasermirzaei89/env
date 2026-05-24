package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint64 extracts uint64 value from env. If not set, returns default value.
func GetUint64(key string, def uint64) uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

// MustGetUint64 extracts uint64 value from env. If not set, it panics.
func MustGetUint64(key string) uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return v
}

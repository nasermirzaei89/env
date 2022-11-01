package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint64 extracts uint64 value from env. if not set, returns default value.
func GetUint64(key string, def uint64) uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize64)
	if err != nil {
		return def
	}

	return v
}

// MustGetUint64 extracts uint64 value from env. if not set, it panics.
func MustGetUint64(key string) uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize64)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return v
}

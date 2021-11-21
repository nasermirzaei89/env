package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt32 extracts int32 value from env. if not set, returns default value.
func GetInt32(key string, def int32) int32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize32)
	if err != nil {
		return def
	}

	return int32(v)
}

// MustGetInt32 extracts int32 value from env. if not set, it panics.
func MustGetInt32(key string) int32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return int32(v)
}

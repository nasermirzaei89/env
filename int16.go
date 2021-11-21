package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt16 extracts int16 value from env. if not set, returns default value.
func GetInt16(key string, def int16) int16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize16)
	if err != nil {
		return def
	}

	return int16(v)
}

// MustGetInt16 extracts int16 value from env. if not set, it panics.
func MustGetInt16(key string) int16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize16)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return int16(v)
}

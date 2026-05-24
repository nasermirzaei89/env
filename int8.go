package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt8 extracts int8 value from env. If not set, returns default value.
func GetInt8(key string, def int8) int8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize8)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return int8(v)
}

// MustGetInt8 extracts int8 value from env. If not set, it panics.
func MustGetInt8(key string) int8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize8)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return int8(v)
}

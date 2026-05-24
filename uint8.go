package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint8 extracts uint8 value from env. If not set, returns default value.
func GetUint8(key string, def uint8) uint8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize8)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return uint8(v)
}

// MustGetUint8 extracts uint8 value from env. If not set, it panics.
func MustGetUint8(key string) uint8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize8)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return uint8(v)
}

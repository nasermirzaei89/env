package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint16 extracts uint16 value from env. If not set, returns default value.
func GetUint16(key string, def uint16) uint16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize16)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return uint16(v)
}

// MustGetUint16 extracts uint16 value from env. If not set, it panics.
func MustGetUint16(key string) uint16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize16)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return uint16(v)
}

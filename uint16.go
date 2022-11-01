package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint16 extracts uint16 value from env. if not set, returns default value.
func GetUint16(key string, def uint16) uint16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize16)
	if err != nil {
		return def
	}

	return uint16(v)
}

// MustGetUint16 extracts uint16 value from env. if not set, it panics.
func MustGetUint16(key string) uint16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize16)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return uint16(v)
}

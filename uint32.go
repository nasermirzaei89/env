package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint32 extracts uint32 value from env. If not set, returns default value.
func GetUint32(key string, def uint32) uint32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return uint32(v)
}

// MustGetUint32 extracts uint32 value from env. If not set, it panics.
func MustGetUint32(key string) uint32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return uint32(v)
}

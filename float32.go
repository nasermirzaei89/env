package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetFloat32 extracts float32 value from env. If not set, returns default value.
func GetFloat32(key string, def float32) float32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseFloat(s, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return float32(v)
}

// MustGetFloat32 extracts float32 value from env. If not set, it panics.
func MustGetFloat32(key string) float32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	v, err := strconv.ParseFloat(s, bitSize32)
	if err != nil {
		panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
	}

	return float32(v)
}

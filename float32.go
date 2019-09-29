package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetFloat32 extracts int value from env
// if not set, returns default value
func GetFloat32(key string, def float32) float32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return def
	}

	return float32(v)
}

// MustGetFloat32 extracts int value from env
// if not set, it panics
func MustGetFloat32(key string) float32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
	}

	return float32(v)
}

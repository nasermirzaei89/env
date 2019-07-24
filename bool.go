package env

import (
	"fmt"
	"os"
)

// GetBool extracts bool value from env
// if not set, returns default value
func GetBool(key string, def bool) bool {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" || s == "1" || s == "true" {
		return true
	}

	return false
}

// MustGetBool extracts bool value from env
// if not set, it panics
func MustGetBool(key string) bool {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' has not been set", key))
	}

	if s == "" || s == "1" || s == "true" {
		return true
	}

	return false
}

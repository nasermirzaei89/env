package env

import (
	"fmt"
	"os"
)

// GetString extracts string value from env. if not set, returns default value.
func GetString(key, def string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return s
}

// MustGetString extracts string value from env. if not set, it panics.
func MustGetString(key string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' has not been set", key))
	}

	return s
}

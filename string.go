package env

import (
	"fmt"
	"os"
)

// GetString extracts string value from env. If not set, returns default value.
func GetString(key, def string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return s
}

// MustGetString extracts string value from env. If not set, it panics.
func MustGetString(key string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	return s
}

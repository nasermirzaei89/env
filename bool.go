package env

import (
	"fmt"
	"os"
)

// GetBool extracts bool value from env. If not set, returns default value.
func GetBool(key string, def bool) bool {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return s == "" || s == "1" || s == "true"
}

// MustGetBool extracts bool value from env. If not set, it panics.
func MustGetBool(key string) bool {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	return s == "" || s == "1" || s == "true"
}

package env

import (
	"fmt"
	"os"
	"strings"
)

// GetStringSlice extracts slice of strings value with format "foo,bar,baz" from env.
// if not set, returns default value.
func GetStringSlice(key string, def []string) []string {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if len(s) == 0 {
		return []string{}
	}

	return strings.Split(s, ",")
}

// MustGetStringSlice extracts slice of strings value with format "foo,bar,baz" from env.
// if not set, it panics.
func MustGetStringSlice(key string) []string {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' has not been set", key))
	}

	if len(s) == 0 {
		return []string{}
	}

	return strings.Split(s, ",")
}

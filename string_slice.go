package env

import (
	"os"
	"strings"
)

// LookupStringSlice extracts slice of string values with format "foo,bar,baz" from env.
// If not set, returns NotSetError.
func LookupStringSlice(key string) ([]string, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if len(s) == 0 {
		return []string{}, nil
	}

	return strings.Split(s, ","), nil
}

// GetStringSlice extracts slice of string values with format "foo,bar,baz" from env.
// If not set, returns default value.
func GetStringSlice(key string, def []string) []string {
	v, err := LookupStringSlice(key)
	if err != nil {
		return def
	}

	return v
}

// MustGetStringSlice extracts slice of string values with format "foo,bar,baz" from env.
// If not set, it panics.
func MustGetStringSlice(key string) []string {
	v, err := LookupStringSlice(key)
	if err != nil {
		panic(err)
	}

	return v
}

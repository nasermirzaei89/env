package env

import (
	"os"
)

// LookupString extracts string value from env. If not set, returns NotSetError.
func LookupString(key string) (string, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return "", NotSetError{Key: key}
	}

	return s, nil
}

// GetString extracts string value from env. If not set, returns default value.
func GetString(key, def string) string {
	v, err := LookupString(key)
	if err != nil {
		return def
	}

	return v
}

// MustGetString extracts string value from env. If not set, it panics.
func MustGetString(key string) string {
	v, err := LookupString(key)
	if err != nil {
		panic(err)
	}

	return v
}

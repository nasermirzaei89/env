package env

import (
	"errors"
	"os"
	"time"
)

// LookupDuration extracts duration value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupDuration(key string) (time.Duration, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := time.ParseDuration(s)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return v, nil
}

// GetDuration extracts duration value from env. If not set, returns default value.
func GetDuration(key string, def time.Duration) time.Duration {
	v, err := LookupDuration(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetDuration extracts duration value from env. If not set, it panics.
func MustGetDuration(key string) time.Duration {
	v, err := LookupDuration(key)
	if err != nil {
		panic(err)
	}

	return v
}

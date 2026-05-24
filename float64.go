package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupFloat64 extracts float64 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupFloat64(key string) (float64, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return v, nil
}

// GetFloat64 extracts float64 value from env. If not set, returns default value.
func GetFloat64(key string, def float64) float64 {
	v, err := LookupFloat64(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetFloat64 extracts float64 value from env. If not set, it panics.
func MustGetFloat64(key string) float64 {
	v, err := LookupFloat64(key)
	if err != nil {
		panic(err)
	}

	return v
}

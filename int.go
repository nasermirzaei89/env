package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupInt extracts int value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupInt(key string) (int, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return v, nil
}

// GetInt extracts int value from env. If not set, returns default value.
func GetInt(key string, def int) int {
	v, err := LookupInt(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt extracts int value from env. If not set, it panics.
func MustGetInt(key string) int {
	v, err := LookupInt(key)
	if err != nil {
		panic(err)
	}

	return v
}

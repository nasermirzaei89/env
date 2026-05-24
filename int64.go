package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupInt64 extracts int64 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupInt64(key string) (int64, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize64)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return v, nil
}

// GetInt64 extracts int64 value from env. If not set, returns default value.
func GetInt64(key string, def int64) int64 {
	v, err := LookupInt64(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt64 extracts int64 value from env. If not set, it panics.
func MustGetInt64(key string) int64 {
	v, err := LookupInt64(key)
	if err != nil {
		panic(err)
	}

	return v
}

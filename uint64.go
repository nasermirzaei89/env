package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupUint64 extracts uint64 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupUint64(key string) (uint64, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize64)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return v, nil
}

// GetUint64 extracts uint64 value from env. If not set, returns default value.
func GetUint64(key string, def uint64) uint64 {
	v, err := LookupUint64(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint64 extracts uint64 value from env. If not set, it panics.
func MustGetUint64(key string) uint64 {
	v, err := LookupUint64(key)
	if err != nil {
		panic(err)
	}

	return v
}

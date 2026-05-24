package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupInt16 extracts int16 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupInt16(key string) (int16, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize16)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return int16(v), nil
}

// GetInt16 extracts int16 value from env. If not set, returns default value.
func GetInt16(key string, def int16) int16 {
	v, err := LookupInt16(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt16 extracts int16 value from env. If not set, it panics.
func MustGetInt16(key string) int16 {
	v, err := LookupInt16(key)
	if err != nil {
		panic(err)
	}

	return v
}

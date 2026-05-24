package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupInt8 extracts int8 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupInt8(key string) (int8, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize8)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return int8(v), nil
}

// GetInt8 extracts int8 value from env. If not set, returns default value.
func GetInt8(key string, def int8) int8 {
	v, err := LookupInt8(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt8 extracts int8 value from env. If not set, it panics.
func MustGetInt8(key string) int8 {
	v, err := LookupInt8(key)
	if err != nil {
		panic(err)
	}

	return v
}

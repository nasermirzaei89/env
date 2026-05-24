package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupInt32 extracts int32 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupInt32(key string) (int32, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseInt(s, decimalBase, bitSize32)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return int32(v), nil
}

// GetInt32 extracts int32 value from env. If not set, returns default value.
func GetInt32(key string, def int32) int32 {
	v, err := LookupInt32(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt32 extracts int32 value from env. If not set, it panics.
func MustGetInt32(key string) int32 {
	v, err := LookupInt32(key)
	if err != nil {
		panic(err)
	}

	return v
}

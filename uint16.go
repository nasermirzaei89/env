package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupUint16 extracts uint16 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupUint16(key string) (uint16, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize16)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return uint16(v), nil
}

// GetUint16 extracts uint16 value from env. If not set, returns default value.
func GetUint16(key string, def uint16) uint16 {
	v, err := LookupUint16(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint16 extracts uint16 value from env. If not set, it panics.
func MustGetUint16(key string) uint16 {
	v, err := LookupUint16(key)
	if err != nil {
		panic(err)
	}

	return v
}

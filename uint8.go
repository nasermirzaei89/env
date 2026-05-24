package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupUint8 extracts uint8 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupUint8(key string) (uint8, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize8)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return uint8(v), nil
}

// GetUint8 extracts uint8 value from env. If not set, returns default value.
func GetUint8(key string, def uint8) uint8 {
	v, err := LookupUint8(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint8 extracts uint8 value from env. If not set, it panics.
func MustGetUint8(key string) uint8 {
	v, err := LookupUint8(key)
	if err != nil {
		panic(err)
	}

	return v
}

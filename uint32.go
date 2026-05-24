package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupUint32 extracts uint32 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupUint32(key string) (uint32, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseUint(s, decimalBase, bitSize32)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return uint32(v), nil
}

// GetUint32 extracts uint32 value from env. If not set, returns default value.
func GetUint32(key string, def uint32) uint32 {
	v, err := LookupUint32(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint32 extracts uint32 value from env. If not set, it panics.
func MustGetUint32(key string) uint32 {
	v, err := LookupUint32(key)
	if err != nil {
		panic(err)
	}

	return v
}

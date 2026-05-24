package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupFloat32 extracts float32 value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupFloat32(key string) (float32, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Key: key}
	}

	v, err := strconv.ParseFloat(s, bitSize32)
	if err != nil {
		return 0, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return float32(v), nil
}

// GetFloat32 extracts float32 value from env. If not set, returns default value.
func GetFloat32(key string, def float32) float32 {
	v, err := LookupFloat32(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetFloat32 extracts float32 value from env. If not set, it panics.
func MustGetFloat32(key string) float32 {
	v, err := LookupFloat32(key)
	if err != nil {
		panic(err)
	}

	return v
}

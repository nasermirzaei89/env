package env

import (
	"errors"
	"os"
	"strconv"
)

// LookupBool extracts bool value from env. If not set, returns NotSetError.
// If the value cannot be parsed as bool, returns InvalidValueError.
func LookupBool(key string) (bool, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return false, NotSetError{Key: key}
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, InvalidValueError{Key: key, Value: s, Err: err}
	}

	return v, nil
}

// GetBool extracts bool value from env. If not set, returns default value.
func GetBool(key string, def bool) bool {
	v, err := LookupBool(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetBool extracts bool value from env. If not set, it panics.
func MustGetBool(key string) bool {
	v, err := LookupBool(key)
	if err != nil {
		panic(err)
	}

	return v
}

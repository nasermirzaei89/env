package env

import "errors"

// LookupUint extracts uint value from env. If not set, returns NotSetError.
// If the value cannot be parsed, returns InvalidValueError.
func LookupUint(key string) (uint, error) {
	v, err := LookupUint64(key)
	if err != nil {
		return 0, err
	}

	return uint(v), nil
}

// GetUint extracts uint value from env. If not set, returns default value.
func GetUint(key string, def uint) uint {
	v, err := LookupUint(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint extracts uint value from env. If not set, it panics.
func MustGetUint(key string) uint {
	v, err := LookupUint(key)
	if err != nil {
		panic(err)
	}

	return v
}

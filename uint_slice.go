package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupUintSlice extracts slice of uint values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupUintSlice(key string) ([]uint, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []uint{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]uint, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize64)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = uint(v)
	}

	return res, nil
}

// GetUintSlice extracts slice of uint values with the format "1,2,3" from env. If not set, returns default value.
func GetUintSlice(key string, def []uint) []uint {
	v, err := LookupUintSlice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUintSlice extracts slice of uint values with the format "1,2,3" from env. If not set, it panics.
func MustGetUintSlice(key string) []uint {
	v, err := LookupUintSlice(key)
	if err != nil {
		panic(err)
	}

	return v
}

package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupUint64Slice extracts slice of uint64 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupUint64Slice(key string) ([]uint64, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []uint64{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]uint64, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize64)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = v
	}

	return res, nil
}

// GetUint64Slice extracts slice of uint64 values with the format "1,2,3" from env. If not set, returns default value.
func GetUint64Slice(key string, def []uint64) []uint64 {
	v, err := LookupUint64Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint64Slice extracts slice of uint64 values with the format "1,2,3" from env. If not set, it panics.
func MustGetUint64Slice(key string) []uint64 {
	v, err := LookupUint64Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

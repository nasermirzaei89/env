package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupInt64Slice extracts slice of int64 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupInt64Slice(key string) ([]int64, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []int64{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]int64, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize64)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = v
	}

	return res, nil
}

// GetInt64Slice extracts slice of int64 values with the format "1,2,3" from env. If not set, returns default value.
func GetInt64Slice(key string, def []int64) []int64 {
	v, err := LookupInt64Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt64Slice extracts slice of int64 values with the format "1,2,3" from env. If not set, it panics.
func MustGetInt64Slice(key string) []int64 {
	v, err := LookupInt64Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

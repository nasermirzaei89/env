package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupInt8Slice extracts slice of int8 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupInt8Slice(key string) ([]int8, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []int8{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]int8, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize8)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = int8(v)
	}

	return res, nil
}

// GetInt8Slice extracts slice of int8 values with the format "1,2,3" from env. If not set, returns default value.
func GetInt8Slice(key string, def []int8) []int8 {
	v, err := LookupInt8Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt8Slice extracts slice of int8 values with the format "1,2,3" from env. If not set, it panics.
func MustGetInt8Slice(key string) []int8 {
	v, err := LookupInt8Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

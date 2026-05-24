package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupInt32Slice extracts slice of int32 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupInt32Slice(key string) ([]int32, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []int32{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]int32, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize32)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = int32(v)
	}

	return res, nil
}

// GetInt32Slice extracts slice of int32 values with the format "1,2,3" from env. If not set, returns default value.
func GetInt32Slice(key string, def []int32) []int32 {
	v, err := LookupInt32Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt32Slice extracts slice of int32 values with the format "1,2,3" from env. If not set, it panics.
func MustGetInt32Slice(key string) []int32 {
	v, err := LookupInt32Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

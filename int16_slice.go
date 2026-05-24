package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupInt16Slice extracts slice of int16 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupInt16Slice(key string) ([]int16, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []int16{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]int16, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize16)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = int16(v)
	}

	return res, nil
}

// GetInt16Slice extracts slice of int16 values with the format "1,2,3" from env. If not set, returns default value.
func GetInt16Slice(key string, def []int16) []int16 {
	v, err := LookupInt16Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetInt16Slice extracts slice of int16 values with the format "1,2,3" from env. If not set, it panics.
func MustGetInt16Slice(key string) []int16 {
	v, err := LookupInt16Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupFloat64Slice extracts slice of float64 values with the format "1.2,2.3,3.4" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupFloat64Slice(key string) ([]float64, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []float64{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]float64, len(ss))

	for i := range ss {
		v, err := strconv.ParseFloat(ss[i], bitSize64)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = v
	}

	return res, nil
}

// GetFloat64Slice extracts slice of float64 values with the format "1.2,2.3,3.4" from env.
// If not set, returns default value.
func GetFloat64Slice(key string, def []float64) []float64 {
	v, err := LookupFloat64Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetFloat64Slice extracts slice of float64 values with the format "1.2,2.3,3.4" from env. If not set, it panics.
func MustGetFloat64Slice(key string) []float64 {
	v, err := LookupFloat64Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

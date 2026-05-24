package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupFloat32Slice extracts slice of float32 values with the format "1.2,2.3,3.4" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupFloat32Slice(key string) ([]float32, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []float32{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]float32, len(ss))

	for i := range ss {
		v, err := strconv.ParseFloat(ss[i], bitSize32)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = float32(v)
	}

	return res, nil
}

// GetFloat32Slice extracts slice of float32 values with the format "1.2,2.3,3.4" from env.
// If not set, returns default value.
func GetFloat32Slice(key string, def []float32) []float32 {
	v, err := LookupFloat32Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetFloat32Slice extracts slice of float32 values with the format "1.2,2.3,3.4" from env. If not set, it panics.
func MustGetFloat32Slice(key string) []float32 {
	v, err := LookupFloat32Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

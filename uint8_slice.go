package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupUint8Slice extracts slice of uint8 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupUint8Slice(key string) ([]uint8, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []uint8{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]uint8, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize8)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = uint8(v)
	}

	return res, nil
}

// GetUint8Slice extracts slice of uint8 values with the format "1,2,3" from env. If not set, returns default value.
func GetUint8Slice(key string, def []uint8) []uint8 {
	v, err := LookupUint8Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint8Slice extracts slice of uint8 values with the format "1,2,3" from env. If not set, it panics.
func MustGetUint8Slice(key string) []uint8 {
	v, err := LookupUint8Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupUint16Slice extracts slice of uint16 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupUint16Slice(key string) ([]uint16, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []uint16{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]uint16, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize16)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = uint16(v)
	}

	return res, nil
}

// GetUint16Slice extracts slice of uint16 values with the format "1,2,3" from env. If not set, returns default value.
func GetUint16Slice(key string, def []uint16) []uint16 {
	v, err := LookupUint16Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint16Slice extracts slice of uint16 values with the format "1,2,3" from env. If not set, it panics.
func MustGetUint16Slice(key string) []uint16 {
	v, err := LookupUint16Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

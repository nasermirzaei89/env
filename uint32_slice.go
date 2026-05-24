package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupUint32Slice extracts slice of uint32 values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupUint32Slice(key string) ([]uint32, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []uint32{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]uint32, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize32)
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = uint32(v)
	}

	return res, nil
}

// GetUint32Slice extracts slice of uint32 values with the format "1,2,3" from env. If not set, returns default value.
func GetUint32Slice(key string, def []uint32) []uint32 {
	v, err := LookupUint32Slice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetUint32Slice extracts slice of uint32 values with the format "1,2,3" from env. If not set, it panics.
func MustGetUint32Slice(key string) []uint32 {
	v, err := LookupUint32Slice(key)
	if err != nil {
		panic(err)
	}

	return v
}

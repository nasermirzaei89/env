package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// LookupIntSlice extracts slice of int values with the format "1,2,3" from env.
// If not set, returns NotSetError. If the value cannot be parsed, returns InvalidValueError.
func LookupIntSlice(key string) ([]int, error) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return nil, NotSetError{Key: key}
	}

	if s == "" {
		return []int{}, nil
	}

	ss := strings.Split(s, ",")

	res := make([]int, len(ss))

	for i := range ss {
		v, err := strconv.Atoi(ss[i])
		if err != nil {
			return nil, InvalidValueError{Key: key, Value: s, Err: err}
		}

		res[i] = v
	}

	return res, nil
}

// GetIntSlice extracts slice of int values with the format "1,2,3" from env. If not set, returns default value.
func GetIntSlice(key string, def []int) []int {
	v, err := LookupIntSlice(key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGetIntSlice extracts slice of int values with the format "1,2,3" from env. If not set, it panics.
func MustGetIntSlice(key string) []int {
	v, err := LookupIntSlice(key)
	if err != nil {
		panic(err)
	}

	return v
}

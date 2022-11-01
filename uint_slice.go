package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUintSlice extracts slice of uint value with the format "1,2,3" from env. if not set, returns default value.
func GetUintSlice(key string, def []uint) []uint {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []uint{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize64)
		if err != nil {
			return def
		}

		res[i] = uint(v)
	}

	return res
}

// MustGetUintSlice extracts slice of uint value with the format "1,2,3" from env. if not set, it panics.
func MustGetUintSlice(key string) []uint {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []uint{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize64)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = uint(v)
	}

	return res
}

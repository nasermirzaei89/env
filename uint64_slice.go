package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUint64Slice extracts slice of uint64 values with the format "1,2,3" from env. If not set, returns default value.
func GetUint64Slice(key string, def []uint64) []uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []uint64{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint64, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize64)
		if err != nil {
			panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
		}

		res[i] = v
	}

	return res
}

// MustGetUint64Slice extracts slice of uint64 values with the format "1,2,3" from env. If not set, it panics.
func MustGetUint64Slice(key string) []uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	if s == "" {
		return []uint64{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint64, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize64)
		if err != nil {
			panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
		}

		res[i] = v
	}

	return res
}

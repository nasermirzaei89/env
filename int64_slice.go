package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetInt64Slice extracts slice of int64 value with the format "1,2,3" from env. if not set, returns default value.
func GetInt64Slice(key string, def []int64) []int64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []int64{}
	}

	ss := strings.Split(s, ",")

	res := make([]int64, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize64)
		if err != nil {
			return def
		}

		res[i] = v
	}

	return res
}

// MustGetInt64Slice extracts slice of int64 value with the format "1,2,3" from env. if not set, it panics.
func MustGetInt64Slice(key string) []int64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []int64{}
	}

	ss := strings.Split(s, ",")

	res := make([]int64, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize64)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = v
	}

	return res
}

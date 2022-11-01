package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetInt8Slice extracts slice of int8 value with the format "1,2,3" from env. if not set, returns default value.
func GetInt8Slice(key string, def []int8) []int8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []int8{}
	}

	ss := strings.Split(s, ",")

	res := make([]int8, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize8)
		if err != nil {
			return def
		}

		res[i] = int8(v)
	}

	return res
}

// MustGetInt8Slice extracts slice of int8 value with the format "1,2,3" from env. if not set, it panics.
func MustGetInt8Slice(key string) []int8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []int8{}
	}

	ss := strings.Split(s, ",")

	res := make([]int8, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize8)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = int8(v)
	}

	return res
}

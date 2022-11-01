package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetInt32Slice extracts slice of int32 value with the format "1,2,3" from env. if not set, returns default value.
func GetInt32Slice(key string, def []int32) []int32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []int32{}
	}

	ss := strings.Split(s, ",")

	res := make([]int32, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize32)
		if err != nil {
			return def
		}

		res[i] = int32(v)
	}

	return res
}

// MustGetInt32Slice extracts slice of int32 value with the format "1,2,3" from env. if not set, it panics.
func MustGetInt32Slice(key string) []int32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []int32{}
	}

	ss := strings.Split(s, ",")

	res := make([]int32, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize32)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = int32(v)
	}

	return res
}

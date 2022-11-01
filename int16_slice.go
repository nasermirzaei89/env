package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetInt16Slice extracts slice of int16 value with the format "1,2,3" from env. if not set, returns default value.
func GetInt16Slice(key string, def []int16) []int16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []int16{}
	}

	ss := strings.Split(s, ",")

	res := make([]int16, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize16)
		if err != nil {
			return def
		}

		res[i] = int16(v)
	}

	return res
}

// MustGetInt16Slice extracts slice of int16 value with the format "1,2,3" from env. if not set, it panics.
func MustGetInt16Slice(key string) []int16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []int16{}
	}

	ss := strings.Split(s, ",")

	res := make([]int16, len(ss))

	for i := range ss {
		v, err := strconv.ParseInt(ss[i], decimalBase, bitSize16)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = int16(v)
	}

	return res
}

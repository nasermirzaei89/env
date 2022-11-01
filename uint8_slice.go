package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUint8Slice extracts slice of uint8 value with the format "1,2,3" from env. if not set, returns default value.
func GetUint8Slice(key string, def []uint8) []uint8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []uint8{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint8, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize8)
		if err != nil {
			return def
		}

		res[i] = uint8(v)
	}

	return res
}

// MustGetUint8Slice extracts slice of uint8 value with the format "1,2,3" from env. if not set, it panics.
func MustGetUint8Slice(key string) []uint8 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []uint8{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint8, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize8)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = uint8(v)
	}

	return res
}

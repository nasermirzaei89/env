package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUint32Slice extracts slice of uint32 value with the format "1,2,3" from env. if not set, returns default value.
func GetUint32Slice(key string, def []uint32) []uint32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []uint32{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint32, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize32)
		if err != nil {
			return def
		}

		res[i] = uint32(v)
	}

	return res
}

// MustGetUint32Slice extracts slice of uint32 value with the format "1,2,3" from env. if not set, it panics.
func MustGetUint32Slice(key string) []uint32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []uint32{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint32, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize32)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = uint32(v)
	}

	return res
}

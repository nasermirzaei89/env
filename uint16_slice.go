package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUint16Slice extracts slice of uint16 value with the format "1,2,3" from env. if not set, returns default value.
func GetUint16Slice(key string, def []uint16) []uint16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []uint16{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint16, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize16)
		if err != nil {
			return def
		}

		res[i] = uint16(v)
	}

	return res
}

// MustGetUint16Slice extracts slice of uint16 value with the format "1,2,3" from env. if not set, it panics.
func MustGetUint16Slice(key string) []uint16 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []uint16{}
	}

	ss := strings.Split(s, ",")

	res := make([]uint16, len(ss))

	for i := range ss {
		v, err := strconv.ParseUint(ss[i], decimalBase, bitSize16)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = uint16(v)
	}

	return res
}

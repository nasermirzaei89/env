package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFloat32Slice extracts slice of float32 value with the format "1.2,2.3,3.4" from env. if not set, returns default value.
func GetFloat32Slice(key string, def []float32) []float32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []float32{}
	}

	ss := strings.Split(s, ",")

	res := make([]float32, len(ss))

	for i := range ss {
		v, err := strconv.ParseFloat(ss[i], bitSize32)
		if err != nil {
			return def
		}

		res[i] = float32(v)
	}

	return res
}

// MustGetFloat32Slice extracts slice of float32 value with the format "1.2,2.3,3.4" from env. if not set, it panics.
func MustGetFloat32Slice(key string) []float32 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []float32{}
	}

	ss := strings.Split(s, ",")

	res := make([]float32, len(ss))

	for i := range ss {
		v, err := strconv.ParseFloat(ss[i], bitSize32)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = float32(v)
	}

	return res
}

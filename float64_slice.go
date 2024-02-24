package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFloat64Slice extracts slice of float64 value with the format "1.2,2.3,3.4" from env. if not set, returns default value.
func GetFloat64Slice(key string, def []float64) []float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []float64{}
	}

	ss := strings.Split(s, ",")

	res := make([]float64, len(ss))

	for i := range ss {
		v, err := strconv.ParseFloat(ss[i], bitSize64)
		if err != nil {
			return def
		}

		res[i] = float64(v)
	}

	return res
}

// MustGetFloat64Slice extracts slice of float64 value with the format "1.2,2.3,3.4" from env. if not set, it panics.
func MustGetFloat64Slice(key string) []float64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable '%s' not set", key))
	}

	if s == "" {
		return []float64{}
	}

	ss := strings.Split(s, ",")

	res := make([]float64, len(ss))

	for i := range ss {
		v, err := strconv.ParseFloat(ss[i], bitSize64)
		if err != nil {
			panic(fmt.Sprintf("invalid environment variable '%s' has been set: %s", key, s))
		}

		res[i] = float64(v)
	}

	return res
}

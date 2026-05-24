package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetIntSlice extracts slice of int values with the format "1,2,3" from env. If not set, returns default value.
func GetIntSlice(key string, def []int) []int {
	s, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	if s == "" {
		return []int{}
	}

	ss := strings.Split(s, ",")

	res := make([]int, len(ss))

	for i := range ss {
		v, err := strconv.Atoi(ss[i])
		if err != nil {
			panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
		}

		res[i] = v
	}

	return res
}

// MustGetIntSlice extracts slice of int values with the format "1,2,3" from env. If not set, it panics.
func MustGetIntSlice(key string) []int {
	s, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}

	if s == "" {
		return []int{}
	}

	ss := strings.Split(s, ",")

	res := make([]int, len(ss))

	for i := range ss {
		v, err := strconv.Atoi(ss[i])
		if err != nil {
			panic(fmt.Sprintf("environment variable %q has an invalid value: %q", key, s))
		}

		res[i] = v
	}

	return res
}

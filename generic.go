package env

import (
	"errors"
	"time"
)

// Type is a constraint that includes all supported types by env package.
type Type interface {
	bool |
		float32 | []float32 | float64 | []float64 |
		int | []int | int8 | []int8 | int16 | []int16 | int32 | []int32 | int64 | []int64 |
		string | []string |
		time.Duration |
		uint | []uint | uint8 | []uint8 | uint16 | []uint16 | uint32 | []uint32 | uint64 | []uint64
}

// Get extracts value from env based on its type. If not set, returns default value.
func Get[T Type](key string, def T) T {
	v, err := Lookup[T](key)
	if err == nil {
		return v
	}

	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return def
	}

	panic(err)
}

// MustGet extracts value from env based on its type. If not set, it panics.
func MustGet[T Type](key string) T {
	v, err := Lookup[T](key)
	if err != nil {
		panic(err)
	}

	return v
}

// Lookup extracts value from env based on its type.
// If not set, returns zero value and NotSetError.
// If the value cannot be parsed, returns zero value and InvalidValueError.
func Lookup[T Type](key string) (res T, err error) {
	var v T

	switch any(v).(type) {
	case bool:
		r, e := LookupBool(key)
		res, err = any(r).(T), e
	case float32:
		r, e := LookupFloat32(key)
		res, err = any(r).(T), e
	case []float32:
		r, e := LookupFloat32Slice(key)
		res, err = any(r).(T), e
	case float64:
		r, e := LookupFloat64(key)
		res, err = any(r).(T), e
	case []float64:
		r, e := LookupFloat64Slice(key)
		res, err = any(r).(T), e
	case int:
		r, e := LookupInt(key)
		res, err = any(r).(T), e
	case []int:
		r, e := LookupIntSlice(key)
		res, err = any(r).(T), e
	case int8:
		r, e := LookupInt8(key)
		res, err = any(r).(T), e
	case []int8:
		r, e := LookupInt8Slice(key)
		res, err = any(r).(T), e
	case int16:
		r, e := LookupInt16(key)
		res, err = any(r).(T), e
	case []int16:
		r, e := LookupInt16Slice(key)
		res, err = any(r).(T), e
	case int32:
		r, e := LookupInt32(key)
		res, err = any(r).(T), e
	case []int32:
		r, e := LookupInt32Slice(key)
		res, err = any(r).(T), e
	case int64:
		r, e := LookupInt64(key)
		res, err = any(r).(T), e
	case []int64:
		r, e := LookupInt64Slice(key)
		res, err = any(r).(T), e
	case string:
		r, e := LookupString(key)
		res, err = any(r).(T), e
	case []string:
		r, e := LookupStringSlice(key)
		res, err = any(r).(T), e
	case time.Duration:
		r, e := LookupDuration(key)
		res, err = any(r).(T), e
	case uint:
		r, e := LookupUint(key)
		res, err = any(r).(T), e
	case []uint:
		r, e := LookupUintSlice(key)
		res, err = any(r).(T), e
	case uint8:
		r, e := LookupUint8(key)
		res, err = any(r).(T), e
	case []uint8:
		r, e := LookupUint8Slice(key)
		res, err = any(r).(T), e
	case uint16:
		r, e := LookupUint16(key)
		res, err = any(r).(T), e
	case []uint16:
		r, e := LookupUint16Slice(key)
		res, err = any(r).(T), e
	case uint32:
		r, e := LookupUint32(key)
		res, err = any(r).(T), e
	case []uint32:
		r, e := LookupUint32Slice(key)
		res, err = any(r).(T), e
	case uint64:
		r, e := LookupUint64(key)
		res, err = any(r).(T), e
	case []uint64:
		r, e := LookupUint64Slice(key)
		res, err = any(r).(T), e
	}

	return //nolint:nakedret
}

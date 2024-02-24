package env

import (
	"fmt"
)

type Type interface {
	bool |
		float32 | []float32 | float64 | []float64 |
		int | []int | int8 | []int8 | int16 | []int16 | int32 | []int32 | int64 | []int64 |
		string | []string |
		uint | []uint | uint8 | []uint8 | uint16 | []uint16 | uint32 | []uint32 | uint64 | []uint64
}

// Get extracts value from env based on its type. if not set, returns default value.
func Get[T Type](key string, def T) T {
	var v T

	switch t := any(v).(type) {
	case bool:
		return any(GetBool(key, any(def).(bool))).(T)
	case float32:
		return any(GetFloat32(key, any(def).(float32))).(T)
	case []float32:
		return any(GetFloat32Slice(key, any(def).([]float32))).(T)
	case float64:
		return any(GetFloat64(key, any(def).(float64))).(T)
	case []float64:
		return any(GetFloat64Slice(key, any(def).([]float64))).(T)
	case int:
		return any(GetInt(key, any(def).(int))).(T)
	case []int:
		return any(GetIntSlice(key, any(def).([]int))).(T)
	case int8:
		return any(GetInt8(key, any(def).(int8))).(T)
	case []int8:
		return any(GetInt8Slice(key, any(def).([]int8))).(T)
	case int16:
		return any(GetInt16(key, any(def).(int16))).(T)
	case []int16:
		return any(GetInt16Slice(key, any(def).([]int16))).(T)
	case int32:
		return any(GetInt32(key, any(def).(int32))).(T)
	case []int32:
		return any(GetInt32Slice(key, any(def).([]int32))).(T)
	case int64:
		return any(GetInt64(key, any(def).(int64))).(T)
	case []int64:
		return any(GetInt64Slice(key, any(def).([]int64))).(T)
	case string:
		return any(GetString(key, any(def).(string))).(T)
	case []string:
		return any(GetStringSlice(key, any(def).([]string))).(T)
	case uint:
		return any(GetUint(key, any(def).(uint))).(T)
	case []uint:
		return any(GetUintSlice(key, any(def).([]uint))).(T)
	case uint8:
		return any(GetUint8(key, any(def).(uint8))).(T)
	case []uint8:
		return any(GetUint8Slice(key, any(def).([]uint8))).(T)
	case uint16:
		return any(GetUint16(key, any(def).(uint16))).(T)
	case []uint16:
		return any(GetUint16Slice(key, any(def).([]uint16))).(T)
	case uint32:
		return any(GetUint32(key, any(def).(uint32))).(T)
	case []uint32:
		return any(GetUint32Slice(key, any(def).([]uint32))).(T)
	case uint64:
		return any(GetUint64(key, any(def).(uint64))).(T)
	case []uint64:
		return any(GetUint64Slice(key, any(def).([]uint64))).(T)
	default:
		panic(fmt.Sprintf("type '%T' is not supported", t))
	}
}

// MustGet extracts string value from env based on its type. if not set, it panics.
func MustGet[T Type](key string) T {
	var v T

	switch t := any(v).(type) {
	case bool:
		return any(MustGetBool(key)).(T)
	case float32:
		return any(MustGetFloat32(key)).(T)
	case []float32:
		return any(MustGetFloat32Slice(key)).(T)
	case float64:
		return any(MustGetFloat64(key)).(T)
	case []float64:
		return any(MustGetFloat64Slice(key)).(T)
	case int:
		return any(MustGetInt(key)).(T)
	case []int:
		return any(MustGetIntSlice(key)).(T)
	case int8:
		return any(MustGetInt8(key)).(T)
	case []int8:
		return any(MustGetInt8Slice(key)).(T)
	case int16:
		return any(MustGetInt16(key)).(T)
	case []int16:
		return any(MustGetInt16Slice(key)).(T)
	case int32:
		return any(MustGetInt32(key)).(T)
	case []int32:
		return any(MustGetInt32Slice(key)).(T)
	case int64:
		return any(MustGetInt64(key)).(T)
	case []int64:
		return any(MustGetInt64Slice(key)).(T)
	case string:
		return any(MustGetString(key)).(T)
	case []string:
		return any(MustGetStringSlice(key)).(T)
	case uint:
		return any(MustGetUint(key)).(T)
	case []uint:
		return any(MustGetUintSlice(key)).(T)
	case uint8:
		return any(MustGetUint8(key)).(T)
	case []uint8:
		return any(MustGetUint8Slice(key)).(T)
	case uint16:
		return any(MustGetUint16(key)).(T)
	case []uint16:
		return any(MustGetUint16Slice(key)).(T)
	case uint32:
		return any(MustGetUint32(key)).(T)
	case []uint32:
		return any(MustGetUint32Slice(key)).(T)
	case uint64:
		return any(MustGetUint64(key)).(T)
	case []uint64:
		return any(MustGetUint64Slice(key)).(T)
	default:
		panic(fmt.Sprintf("type '%T' is not supported", t))
	}
}

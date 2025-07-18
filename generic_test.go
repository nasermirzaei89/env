package env_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGet(t *testing.T) {
	// bool
	{
		t.Run("GetAbsentBoolWithDefault", func(t *testing.T) {
			res := env.Get("V1", true)
			assertTrue(t, res)
		})

		t.Run("GetPresentBoolWithDefault", func(t *testing.T) {
			t.Setenv("V1", "false")

			res := env.Get("V1", true)
			assertFalse(t, res)
		})

		t.Run("GetZeroAsBoolWithDefault", func(t *testing.T) {
			t.Setenv("V1", "0")

			res := env.Get("V1", true)
			assertFalse(t, res)
		})

		t.Run("GetTrueStringAsBoolWithDefault", func(t *testing.T) {
			t.Setenv("V1", "true")

			res := env.Get("V1", false)
			assertTrue(t, res)
		})

		t.Run("GetEmptyAsBoolWithDefault", func(t *testing.T) {
			t.Setenv("V1", "")

			res := env.Get("V1", false)
			assertTrue(t, res)
		})

		t.Run("GetOneAsBoolWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1")

			res := env.Get("V1", false)
			assertTrue(t, res)
		})
	}

	// float32
	{
		def := float32(12.5)

		t.Run("GetAbsentFloat32WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidFloat32WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)

			assertEqual(t, def, res)
		})

		t.Run("GetValidFloat32WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14.5")

			res := env.Get("V1", def)

			assertEqual(t, float32(14.5), res)
		})
	}

	// []float32
	{
		def := []float32{21.2, 22.3}

		t.Run("GetAbsentFloat32SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidFloat32SliceWithDefault", func(t *testing.T) {
			expected := []float32{31.02, 32.33, 33.33}

			t.Setenv("V1", "31.02,32.33,33.33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidFloat32SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1.2,2.3,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyFloat32SliceWithDefault", func(t *testing.T) {
			expected := make([]float32, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// float64
	{
		def := 12.5

		t.Run("GetAbsentFloat64WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidFloat64WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidFloat64WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14.5")

			res := env.Get("V1", def)
			assertEqual(t, 14.5, res)
		})
	}

	// []float64
	{
		def := []float64{21.2, 22.3}

		t.Run("GetAbsentFloat64SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidFloat64SliceWithDefault", func(t *testing.T) {
			expected := []float64{31.02, 32.33, 33.33}

			t.Setenv("V1", "31.02,32.33,33.33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidFloat64SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1.2,2.3,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyFloat64SliceWithDefault", func(t *testing.T) {
			expected := make([]float64, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// int
	{
		def := 12

		t.Run("GetAbsentIntWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidIntWithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidIntWithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, 14, res)
		})
	}

	// []int
	{
		def := []int{21, 22}

		t.Run("GetAbsentIntSliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidIntSliceWithDefault", func(t *testing.T) {
			expected := []int{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidIntSliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyIntSliceWithDefault", func(t *testing.T) {
			expected := make([]int, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// int8
	{
		def := int8(12)

		t.Run("GetAbsentInt8WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidInt8WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidInt8WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, int8(14), res)
		})
	}

	// []int8
	{
		t.Run("GetAbsentInt8SliceWithDefault", func(t *testing.T) {
			def := []int8{21, 22}

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidInt8SliceWithDefault", func(t *testing.T) {
			def := []int8{21, 22}
			expected := []int8{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidInt8SliceWithDefault", func(t *testing.T) {
			def := []int8{21, 22}

			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetInvalidInt8SliceWithDefault2", func(t *testing.T) {
			def := []int8{21, 22}

			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyInt8SliceWithDefault", func(t *testing.T) {
			def := []int8{21, 22}
			expected := make([]int8, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// int16
	{
		def := int16(12)

		t.Run("GetAbsentInt16WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidInt16WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidInt16WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, int16(14), res)
		})
	}

	// []int16
	{
		def := []int16{21, 22}

		t.Run("GetAbsentInt16SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidInt16SliceWithDefault", func(t *testing.T) {
			expected := []int16{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidInt16SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyInt16SliceWithDefault", func(t *testing.T) {
			expected := make([]int16, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// int32
	{
		def := int32(12)

		t.Run("GetAbsentInt32WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidInt32WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)

			assertEqual(t, def, res)
		})

		t.Run("GetValidInt32WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)

			assertEqual(t, int32(14), res)
		})
	}

	// []int32
	{
		def := []int32{21, 22}

		t.Run("GetAbsentInt32SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidInt32SliceWithDefault", func(t *testing.T) {
			expected := []int32{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidInt32SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyInt32SliceWithDefault", func(t *testing.T) {
			expected := make([]int32, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// int64
	{
		def := int64(12)

		t.Run("GetAbsentInt64WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidInt64WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidInt64WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, int64(14), res)
		})
	}

	// []int64
	{
		def := []int64{21, 22}

		t.Run("GetAbsentInt64SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidInt64SliceWithDefault", func(t *testing.T) {
			expected := []int64{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidInt64SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyInt64SliceWithDefault", func(t *testing.T) {
			expected := make([]int64, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// string
	{
		def := "v1_default"

		t.Run("GetAbsentStringWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidStringWithDefault", func(t *testing.T) {
			t.Setenv("V1", "val")

			res := env.Get("V1", def)
			assertEqual(t, "val", res)
		})
	}

	// []string
	{
		def := []string{"foo", "bar"}

		t.Run("GetAbsentStringSliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidStringSliceWithDefault", func(t *testing.T) {
			expected := []string{"foo", "bar", "baz"}
			t.Setenv("V1", strings.Join(expected, ","))

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetEmptyStringSliceWithDefault", func(t *testing.T) {
			expected := make([]string, 0)
			t.Setenv("V1", strings.Join(expected, ","))

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// uint
	{
		def := uint(12)

		t.Run("GetAbsentUIntWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidUIntWithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidUIntWithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, uint(14), res)
		})
	}

	// []uint
	{
		def := []uint{21, 22}

		t.Run("GetAbsentUIntSliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidUIntSliceWithDefault", func(t *testing.T) {
			expected := []uint{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidUIntSliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyUIntSliceWithDefault", func(t *testing.T) {
			expected := make([]uint, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// uint8
	{
		def := uint8(12)

		t.Run("GetAbsentUInt8WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidUInt8WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidUInt8WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, uint8(14), res)
		})
	}

	// []uint8
	{
		def := []uint8{21, 22}

		t.Run("GetAbsentUInt8SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidUInt8SliceWithDefault", func(t *testing.T) {
			expected := []uint8{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidUInt8SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetInvalidUInt8SliceWithDefault2", func(t *testing.T) {
			expected := make([]uint8, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// uint16
	{
		def := uint16(12)

		t.Run("GetAbsentUInt16WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidUInt16WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidUInt16WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, uint16(14), res)
		})
	}

	// []uint16
	{
		def := []uint16{21, 22}

		t.Run("GetAbsentUInt16SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidUInt16SliceWithDefault", func(t *testing.T) {
			expected := []uint16{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidUInt16SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyUInt16SliceWithDefault", func(t *testing.T) {
			expected := make([]uint16, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// uint32
	{
		def := uint32(12)

		t.Run("GetAbsentUInt32WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidUInt32WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidUInt32WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, uint32(14), res)
		})
	}

	// []uint32
	{
		def := []uint32{21, 22}

		t.Run("GetAbsentUInt32SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidUInt32SliceWithDefault", func(t *testing.T) {
			expected := []uint32{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidUInt32SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyUInt32SliceWithDefault", func(t *testing.T) {
			expected := make([]uint32, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}

	// uint64
	{
		def := uint64(12)

		t.Run("GetAbsentUInt64WithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetInvalidUInt64WithDefault", func(t *testing.T) {
			t.Setenv("V1", "invalid")

			res := env.Get("V1", def)
			assertEqual(t, def, res)
		})

		t.Run("GetValidUInt64WithDefault", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.Get("V1", def)
			assertEqual(t, uint64(14), res)
		})
	}

	// []uint64
	{
		def := []uint64{21, 22}

		t.Run("GetAbsentUInt64SliceWithDefault", func(t *testing.T) {
			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetValidUInt64SliceWithDefault", func(t *testing.T) {
			expected := []uint64{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})

		t.Run("GetInvalidUInt64SliceWithDefault", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			res := env.Get("V1", def)
			assertEqualSlices(t, def, res)
		})

		t.Run("GetEmptyUInt64SliceWithDefault", func(t *testing.T) {
			expected := make([]uint64, 0)

			t.Setenv("V1", "")

			res := env.Get("V1", def)
			assertEqualSlices(t, expected, res)
		})
	}
}

func ExampleGet_string() {
	_ = os.Setenv("V1", "val")

	fmt.Println(env.Get("V1", "default"))
	// Output: val
}

func ExampleGet_stringSlice() {
	_ = os.Setenv("V1", "foo,bar,baz")

	fmt.Println(env.Get("V1", []string{"default"}))
	// Output: [foo bar baz]
}

func ExampleGet_int() {
	_ = os.Setenv("V1", "14")

	fmt.Println(env.Get("V1", 12))
	// Output: 14
}

func ExampleGet_intSlice() {
	_ = os.Setenv("V1", "31,32,33")

	fmt.Println(env.Get("V1", []int{21}))
	// Output: [31 32 33]
}

func ExampleGet_bool() {
	_ = os.Setenv("V1", "true")

	fmt.Println(env.Get("V1", false))
	// Output: true
}

func TestMustGet(t *testing.T) {
	// bool
	{
		t.Run("MustGetAbsentBool", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[bool]("V1")
			})
		})

		t.Run("MustGetFalseStringAsBool", func(t *testing.T) {
			t.Setenv("V1", "false")

			res := env.MustGet[bool]("V1")
			assertFalse(t, res)
		})

		t.Run("MustGetZeroAsBool", func(t *testing.T) {
			t.Setenv("V1", "0")

			res := env.MustGet[bool]("V1")
			assertFalse(t, res)
		})

		t.Run("MustGetTrueStringAsBool", func(t *testing.T) {
			t.Setenv("V1", "true")

			res := env.MustGet[bool]("V1")
			assertTrue(t, res)
		})

		t.Run("MustGetOneAsBool", func(t *testing.T) {
			t.Setenv("V1", "1")

			res := env.MustGet[bool]("V1")
			assertTrue(t, res)
		})
	}

	// float32
	{
		t.Run("MustGetAbsentFloat32", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[float32]("V1")
			})
		})

		t.Run("MustGetInvalidFloat32", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[float32]("V1")
			})
		})

		t.Run("MustGetValidFloat32", func(t *testing.T) {
			t.Setenv("V1", "14.5")

			res := env.MustGet[float32]("V1")
			assertEqual(t, float32(14.5), res)
		})
	}

	// []float32
	{
		t.Run("MustGetAbsentFloat32Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]float32]("V1")
			})
		})

		t.Run("MustGetValidFloat32Slice", func(t *testing.T) {
			expected := []float32{31.02, 32.33, 33.33}

			t.Setenv("V1", "31.02,32.33,33.33")

			res := env.MustGet[[]float32]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidFloat32Slice", func(t *testing.T) {
			t.Setenv("V1", "1.2,2.3,Three")

			assertPanics(t, func() {
				env.MustGet[[]float32]("V1")
			})
		})

		t.Run("MustGetEmptyFloat32Slice", func(t *testing.T) {
			expected := make([]float32, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]float32]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// float64
	{
		t.Run("MustGetAbsentFloat64", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[float64]("V1")
			})
		})

		t.Run("MustGetInvalidFloat64", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")

				env.MustGet[float64]("V1")
			})
		})

		t.Run("MustGetValidFloat64", func(t *testing.T) {
			t.Setenv("V1", "14.5")

			res := env.MustGet[float64]("V1")
			assertEqual(t, 14.5, res)
		})
	}

	// []float64
	{
		t.Run("MustGetAbsentFloat64Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]float64]("V1")
			})
		})

		t.Run("MustGetValidFloat64Slice", func(t *testing.T) {
			expected := []float64{31.02, 32.33, 33.33}

			t.Setenv("V1", "31.02,32.33,33.33")

			res := env.MustGet[[]float64]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidFloat64Slice", func(t *testing.T) {
			t.Setenv("V1", "1.2,2.3,Three")

			assertPanics(t, func() {
				env.MustGet[[]float64]("V1")
			})
		})

		t.Run("MustGetEmptyFloat64Slice", func(t *testing.T) {
			expected := make([]float64, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]float64]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// int
	{
		t.Run("MustGetAbsentInt", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[int]("V1")
			})
		})

		t.Run("MustGetInvalidInt", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[int]("V1")
			})
		})

		t.Run("MustGetValidInt", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[int]("V1")
			assertEqual(t, 14, res)
		})
	}

	// []int
	{
		t.Run("MustGetAbsentIntSlice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]int]("V1")
			})
		})

		t.Run("MustGetValidIntSlice", func(t *testing.T) {
			expected := []int{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]int]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidIntSlice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]int]("V1")
			})
		})

		t.Run("MustGetEmptyIntSlice", func(t *testing.T) {
			expected := make([]int, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]int]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// int8
	{
		t.Run("MustGetAbsentInt8", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[int8]("V1")
			})
		})

		t.Run("MustGetInvalidInt8", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[int8]("V1")
			})
		})

		t.Run("MustGetValidInt8", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[int8]("V1")
			assertEqual(t, int8(14), res)
		})
	}

	// []int8
	{
		t.Run("MustGetAbsentInt8Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]int8]("V1")
			})
		})

		t.Run("MustGetValidInt8Slice", func(t *testing.T) {
			expected := []int8{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]int8]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidInt8Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]int8]("V1")
			})
		})

		t.Run("MustGetEmptyInt8Slice", func(t *testing.T) {
			expected := make([]int8, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]int8]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// int16
	{
		t.Run("MustGetAbsentInt16", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[int16]("V1")
			})
		})

		t.Run("MustGetInvalidInt16", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")

				env.MustGet[int16]("V1")
			})
		})

		t.Run("MustGetValidInt16", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[int16]("V1")
			assertEqual(t, int16(14), res)
		})
	}

	// []int16
	{
		t.Run("MustGetAbsentInt16Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]int16]("V1")
			})
		})

		t.Run("MustGetValidInt16Slice", func(t *testing.T) {
			expected := []int16{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]int16]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidInt16Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]int16]("V1")
			})
		})

		t.Run("MustGetEmptyInt16Slice", func(t *testing.T) {
			expected := make([]int16, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]int16]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// int32
	{
		t.Run("MustGetAbsentInt32", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[int32]("V1")
			})
		})

		t.Run("MustGetInvalidInt32", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")

				env.MustGet[int32]("V1")
			})
		})

		t.Run("MustGetValidInt32", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[int32]("V1")
			assertEqual(t, int32(14), res)
		})
	}

	// []int32
	{
		t.Run("MustGetAbsentInt32Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]int32]("V1")
			})
		})

		t.Run("MustGetValidInt32Slice", func(t *testing.T) {
			expected := []int32{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]int32]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidInt32Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]int32]("V1")
			})
		})

		t.Run("MustGetEmptyInt32Slice", func(t *testing.T) {
			expected := make([]int32, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]int32]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// int64
	{
		t.Run("MustGetAbsentInt64", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[int64]("V1")
			})
		})

		t.Run("MustGetInvalidInt64", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[int64]("V1")
			})
		})

		t.Run("MustGetValidInt64", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[int64]("V1")
			assertEqual(t, int64(14), res)
		})
	}

	// []int64
	{
		t.Run("MustGetAbsentInt64Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]int64]("V1")
			})
		})

		t.Run("MustGetValidInt64Slice", func(t *testing.T) {
			expected := []int64{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]int64]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidInt64Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]int64]("V1")
			})
		})

		t.Run("MustGetEmptyInt64Slice", func(t *testing.T) {
			expected := make([]int64, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]int64]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// string
	{
		t.Run("MustGetAbsentString", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[string]("V1")
			})
		})

		t.Run("MustGetValidString", func(t *testing.T) {
			t.Setenv("V1", "val")

			res := env.MustGet[string]("V1")
			assertEqual(t, "val", res)
		})
	}

	// []string
	{
		t.Run("MustGetAbsentStringSlice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]string]("V1")
			})
		})

		t.Run("MustGetValidStringSlice", func(t *testing.T) {
			expected := []string{"foo", "bar", "baz"}
			t.Setenv("V1", strings.Join(expected, ","))

			res := env.MustGet[[]string]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetEmptyStringSlice", func(t *testing.T) {
			expected := make([]string, 0)
			t.Setenv("V1", strings.Join(expected, ","))

			res := env.MustGet[[]string]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// uint
	{
		t.Run("MustGetAbsentUInt", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[uint]("V1")
			})
		})

		t.Run("MustGetInvalidUInt", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[uint]("V1")
			})
		})

		t.Run("MustGetValidUInt", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[uint]("V1")
			assertEqual(t, uint(14), res)
		})
	}

	// []uint
	{
		t.Run("MustGetAbsentUIntSlice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]uint]("V1")
			})
		})

		t.Run("MustGetValidUIntSlice", func(t *testing.T) {
			expected := []uint{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]uint]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidUIntSlice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]uint]("V1")
			})
		})

		t.Run("MustGetEmptyUIntSlice", func(t *testing.T) {
			expected := make([]uint, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]uint]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// uint8
	{
		t.Run("MustGetAbsentUInt8", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[uint8]("V1")
			})
		})

		t.Run("MustGetInvalidUInt8", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[uint8]("V1")
			})
		})

		t.Run("MustGetValidUInt8", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[uint8]("V1")
			assertEqual(t, uint8(14), res)
		})
	}

	// []uint8
	{
		t.Run("MustGetAbsentUInt8Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]uint8]("V1")
			})
		})

		t.Run("MustGetValidUInt8Slice", func(t *testing.T) {
			expected := []uint8{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]uint8]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidUInt8Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]uint8]("V1")
			})
		})

		t.Run("MustGetEmptyUInt8Slice", func(t *testing.T) {
			expected := make([]uint8, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]uint8]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// uint16
	{
		t.Run("MustGetAbsentUInt16", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[uint16]("V1")
			})
		})

		t.Run("MustGetInvalidUInt16", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[uint16]("V1")
			})
		})

		t.Run("MustGetValidUInt16", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[uint16]("V1")
			assertEqual(t, uint16(14), res)
		})
	}

	// []uint16
	{
		t.Run("MustGetAbsentUInt16Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]uint16]("V1")
			})
		})

		t.Run("MustGetValidUInt16Slice", func(t *testing.T) {
			expected := []uint16{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]uint16]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidUInt16Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]uint16]("V1")
			})
		})

		t.Run("MustGetEmptyUInt16Slice", func(t *testing.T) {
			expected := make([]uint16, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]uint16]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// uint32
	{
		t.Run("MustGetAbsentUInt32", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[uint32]("V1")
			})
		})

		t.Run("MustGetInvalidUInt32", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[uint32]("V1")
			})
		})

		t.Run("MustGetValidUInt32", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[uint32]("V1")
			assertEqual(t, uint32(14), res)
		})
	}

	// []uint32
	{
		t.Run("MustGetAbsentUInt32Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]uint32]("V1")
			})
		})

		t.Run("MustGetValidUInt32Slice", func(t *testing.T) {
			expected := []uint32{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]uint32]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidUInt32Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]uint32]("V1")
			})
		})

		t.Run("MustGetEmptyUInt32Slice", func(t *testing.T) {
			expected := make([]uint32, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]uint32]("V1")
			assertEqualSlices(t, expected, res)
		})
	}

	// uint64
	{
		t.Run("MustGetAbsentUInt64", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[uint64]("V1")
			})
		})

		t.Run("MustGetInvalidUInt64", func(t *testing.T) {
			assertPanics(t, func() {
				t.Setenv("V1", "invalid")
				env.MustGet[uint64]("V1")
			})
		})

		t.Run("MustGetValidUInt64", func(t *testing.T) {
			t.Setenv("V1", "14")

			res := env.MustGet[uint64]("V1")
			assertEqual(t, uint64(14), res)
		})
	}

	// []uint64
	{
		t.Run("MustGetAbsentUInt64Slice", func(t *testing.T) {
			assertPanics(t, func() {
				env.MustGet[[]uint64]("V1")
			})
		})

		t.Run("MustGetValidUInt64Slice", func(t *testing.T) {
			expected := []uint64{31, 32, 33}

			t.Setenv("V1", "31,32,33")

			res := env.MustGet[[]uint64]("V1")
			assertEqualSlices(t, expected, res)
		})

		t.Run("MustGetInvalidUInt64Slice", func(t *testing.T) {
			t.Setenv("V1", "1,2,Three")

			assertPanics(t, func() {
				env.MustGet[[]uint64]("V1")
			})
		})

		t.Run("MustGetEmptyUInt64Slice", func(t *testing.T) {
			expected := make([]uint64, 0)

			t.Setenv("V1", "")

			res := env.MustGet[[]uint64]("V1")
			assertEqualSlices(t, expected, res)
		})
	}
}

func ExampleMustGet_string() {
	_ = os.Setenv("V1", "val")

	fmt.Println(env.MustGet[string]("V1"))
	// Output: val
}

func ExampleMustGet_stringSlice() {
	_ = os.Setenv("V1", "foo,bar,baz")

	fmt.Println(env.MustGet[[]string]("V1"))
	// Output: [foo bar baz]
}

func ExampleMustGet_int() {
	_ = os.Setenv("V1", "14")

	fmt.Println(env.MustGet[int]("V1"))
	// Output: 14
}

func ExampleMustGet_intSlice() {
	_ = os.Setenv("V1", "31,32,33")

	fmt.Println(env.MustGet[[]int]("V1"))
	// Output: [31 32 33]
}

func ExampleMustGet_bool() {
	_ = os.Setenv("V1", "true")

	fmt.Println(env.MustGet[bool]("V1"))
	// Output: true
}

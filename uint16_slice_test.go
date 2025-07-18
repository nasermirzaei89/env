package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint16Slice(t *testing.T) {
	def := []uint16{21, 22}

	t.Run("GetAbsentUInt16SliceWithDefault", func(t *testing.T) {
		res := env.GetUint16Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidUInt16SliceWithDefault", func(t *testing.T) {
		expected := []uint16{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetUint16Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidUInt16SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetUint16Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetEmptyUInt16SliceWithDefault", func(t *testing.T) {
		expected := make([]uint16, 0)

		t.Setenv("V1", "")

		res := env.GetUint16Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetUint16Slice(t *testing.T) {
	t.Run("MustGetAbsentUInt16Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint16Slice("V1")
		})
	})

	t.Run("MustGetValidUInt16Slice", func(t *testing.T) {
		expected := []uint16{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetUint16Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidUInt16Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetUint16Slice("V1")
		})
	})

	t.Run("MustGetEmptyUInt16Slice", func(t *testing.T) {
		expected := make([]uint16, 0)

		t.Setenv("V1", "")

		res := env.MustGetUint16Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

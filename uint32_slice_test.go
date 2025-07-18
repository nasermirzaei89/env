package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint32Slice(t *testing.T) {
	def := []uint32{21, 22}

	t.Run("GetAbsentUInt32SliceWithDefault", func(t *testing.T) {
		res := env.GetUint32Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidUInt32SliceWithDefault", func(t *testing.T) {
		expected := []uint32{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetUint32Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidUInt32SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetUint32Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetEmptyUInt32SliceWithDefault", func(t *testing.T) {
		expected := make([]uint32, 0)

		t.Setenv("V1", "")

		res := env.GetUint32Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetUint32Slice(t *testing.T) {
	t.Run("MustGetAbsentUInt32Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint32Slice("V1")
		})
	})

	t.Run("MustGetValidUInt32Slice", func(t *testing.T) {
		expected := []uint32{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetUint32Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidUInt32Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetUint32Slice("V1")
		})
	})

	t.Run("MustGetEmptyUInt32Slice", func(t *testing.T) {
		expected := make([]uint32, 0)

		t.Setenv("V1", "")

		res := env.MustGetUint32Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

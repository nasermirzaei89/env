package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint64Slice(t *testing.T) {
	def := []uint64{21, 22}

	t.Run("GetAbsentUInt64SliceWithDefault", func(t *testing.T) {
		res := env.GetUint64Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidUInt64SliceWithDefault", func(t *testing.T) {
		expected := []uint64{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetUint64Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidUInt64SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetUint64Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetEmptyUInt64SliceWithDefault", func(t *testing.T) {
		expected := make([]uint64, 0)

		t.Setenv("V1", "")

		res := env.GetUint64Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetUint64Slice(t *testing.T) {
	t.Run("MustGetAbsentUInt64Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint64Slice("V1")
		})
	})

	t.Run("MustGetValidUInt64Slice", func(t *testing.T) {
		expected := []uint64{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetUint64Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidUInt64Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetUint64Slice("V1")
		})
	})

	t.Run("MustGetEmptyUInt64Slice", func(t *testing.T) {
		expected := make([]uint64, 0)

		t.Setenv("V1", "")

		res := env.MustGetUint64Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

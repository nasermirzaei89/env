package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt32Slice(t *testing.T) {
	def := []int32{21, 22}

	t.Run("GetAbsentInt32SliceWithDefault", func(t *testing.T) {
		res := env.GetInt32Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidInt32SliceWithDefault", func(t *testing.T) {
		expected := []int32{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetInt32Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidInt32SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetInt32Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetEmptyInt32SliceWithDefault", func(t *testing.T) {
		expected := make([]int32, 0)

		t.Setenv("V1", "")

		res := env.GetInt32Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetInt32Slice(t *testing.T) {
	t.Run("MustGetAbsentInt32Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt32Slice("V1")
		})
	})

	t.Run("MustGetValidInt32Slice", func(t *testing.T) {
		expected := []int32{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetInt32Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidInt32Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetInt32Slice("V1")
		})
	})

	t.Run("MustGetEmptyInt32Slice", func(t *testing.T) {
		expected := make([]int32, 0)

		t.Setenv("V1", "")

		res := env.MustGetInt32Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

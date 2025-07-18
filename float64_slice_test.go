package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetFloat64Slice(t *testing.T) {
	def := []float64{21.2, 22.3}

	t.Run("GetAbsentFloat64SliceWithDefault", func(t *testing.T) {
		res := env.GetFloat64Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidFloat64SliceWithDefault", func(t *testing.T) {
		expected := []float64{31.02, 32.33, 33.33}

		t.Setenv("V1", "31.02,32.33,33.33")

		res := env.GetFloat64Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidFloat64SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1.2,2.3,Three")

		res := env.GetFloat64Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetEmptyFloat64SliceWithDefault", func(t *testing.T) {
		expected := make([]float64, 0)

		t.Setenv("V1", "")

		res := env.GetFloat64Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetFloat64Slice(t *testing.T) {
	t.Run("MustGetAbsentFloat64Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetFloat64Slice("V1")
		})
	})

	t.Run("MustGetValidFloat64Slice", func(t *testing.T) {
		expected := []float64{31.02, 32.33, 33.33}

		t.Setenv("V1", "31.02,32.33,33.33")

		res := env.MustGetFloat64Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidFloat64Slice", func(t *testing.T) {
		t.Setenv("V1", "1.2,2.3,Three")

		assertPanics(t, func() {
			env.MustGetFloat64Slice("V1")
		})
	})

	t.Run("MustGetEmptyFloat64Slice", func(t *testing.T) {
		expected := make([]float64, 0)

		t.Setenv("V1", "")

		res := env.MustGetFloat64Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

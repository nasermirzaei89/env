package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetFloat32Slice(t *testing.T) {
	t.Run("GetAbsentFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}

		res := env.GetFloat32Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}
		expected := []float32{31.02, 32.33, 33.33}

		t.Setenv("V1", "31.02,32.33,33.33")

		res := env.GetFloat32Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}

		t.Setenv("V1", "1.2,2.3,Three")

		assertPanics(t, func() {
			env.GetFloat32Slice("V1", def)
		})
	})

	t.Run("GetEmptyFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}
		expected := make([]float32, 0)

		t.Setenv("V1", "")

		res := env.GetFloat32Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetFloat32Slice(t *testing.T) {
	t.Run("MustGetAbsentFloat32Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetFloat32Slice("V1")
		})
	})

	t.Run("MustGetValidFloat32Slice", func(t *testing.T) {
		expected := []float32{31.02, 32.33, 33.33}

		t.Setenv("V1", "31.02,32.33,33.33")

		res := env.MustGetFloat32Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidFloat32Slice", func(t *testing.T) {
		t.Setenv("V1", "1.2,2.3,Three")

		assertPanics(t, func() {
			env.MustGetFloat32Slice("V1")
		})
	})

	t.Run("MustGetEmptyFloat32Slice", func(t *testing.T) {
		expected := make([]float32, 0)

		t.Setenv("V1", "")

		res := env.MustGetFloat32Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

func TestLookupFloat32Slice(t *testing.T) {
	t.Run("LookupAbsentFloat32Slice", func(t *testing.T) {
		_, err := env.LookupFloat32Slice("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidFloat32Slice", func(t *testing.T) {
		t.Setenv("V1", "1.1,2.2,3.3")

		res, err := env.LookupFloat32Slice("V1")
		assertNoError(t, err)
		assertEqualSlices(t, []float32{1.1, 2.2, 3.3}, res)
	})

	t.Run("LookupInvalidFloat32Slice", func(t *testing.T) {
		t.Setenv("V1", "1.1,2.2,Three")

		_, err := env.LookupFloat32Slice("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
	})
}

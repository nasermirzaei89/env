package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt16Slice(t *testing.T) {
	def := []int16{21, 22}

	t.Run("GetAbsentInt16SliceWithDefault", func(t *testing.T) {
		res := env.GetInt16Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidInt16SliceWithDefault", func(t *testing.T) {
		expected := []int16{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetInt16Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidInt16SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.GetInt16Slice("V1", def)
		})
	})

	t.Run("GetEmptyInt16SliceWithDefault", func(t *testing.T) {
		expected := make([]int16, 0)

		t.Setenv("V1", "")

		res := env.GetInt16Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetInt16Slice(t *testing.T) {
	t.Run("MustGetAbsentInt16Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt16Slice("V1")
		})
	})

	t.Run("MustGetValidInt16Slice", func(t *testing.T) {
		expected := []int16{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetInt16Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidInt16Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetInt16Slice("V1")
		})
	})

	t.Run("MustGetEmptyInt16Slice", func(t *testing.T) {
		expected := make([]int16, 0)

		t.Setenv("V1", "")

		res := env.MustGetInt16Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

func TestLookupInt16Slice(t *testing.T) {
	t.Run("LookupAbsentInt16Slice", func(t *testing.T) {
		_, err := env.LookupInt16Slice("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt16Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,3")

		res, err := env.LookupInt16Slice("V1")
		assertNoError(t, err)
		assertEqualSlices(t, []int16{1, 2, 3}, res)
	})

	t.Run("LookupInvalidInt16Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		_, err := env.LookupInt16Slice("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
	})
}

package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetIntSlice(t *testing.T) {
	def := []int{21, 22}

	t.Run("GetAbsentIntSliceWithDefault", func(t *testing.T) {
		res := env.GetIntSlice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidIntSliceWithDefault", func(t *testing.T) {
		expected := []int{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetIntSlice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidIntSliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.GetIntSlice("V1", def)
		})
	})

	t.Run("GetEmptyIntSliceWithDefault", func(t *testing.T) {
		expected := make([]int, 0)

		t.Setenv("V1", "")

		res := env.GetIntSlice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetIntSlice(t *testing.T) {
	t.Run("MustGetAbsentIntSlice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetIntSlice("V1")
		})
	})

	t.Run("MustGetValidIntSlice", func(t *testing.T) {
		expected := []int{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetIntSlice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidIntSlice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetIntSlice("V1")
		})
	})

	t.Run("MustGetEmptyIntSlice", func(t *testing.T) {
		expected := make([]int, 0)

		t.Setenv("V1", "")

		res := env.MustGetIntSlice("V1")
		assertEqualSlices(t, expected, res)
	})
}

func TestLookupIntSlice(t *testing.T) {
	t.Run("LookupAbsentIntSlice", func(t *testing.T) {
		_, err := env.LookupIntSlice("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidIntSlice", func(t *testing.T) {
		t.Setenv("V1", "31,32,33")

		res, err := env.LookupIntSlice("V1")
		assertNoError(t, err)
		assertEqualSlices(t, []int{31, 32, 33}, res)
	})

	t.Run("LookupInvalidIntSlice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		_, err := env.LookupIntSlice("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
	})
}

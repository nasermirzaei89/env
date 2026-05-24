package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt8Slice(t *testing.T) {
	t.Run("GetAbsentInt8SliceWithDefault", func(t *testing.T) {
		def := []int8{21, 22}

		res := env.GetInt8Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidInt8SliceWithDefault", func(t *testing.T) {
		def := []int8{21, 22}
		expected := []int8{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetInt8Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidInt8SliceWithDefault", func(t *testing.T) {
		def := []int8{21, 22}

		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetInt8Slice("V1", def)
		})
	})

	t.Run("GetInvalidInt8SliceWithDefault2", func(t *testing.T) {
		def := []int8{21, 22}

		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.GetInt8Slice("V1", def)
		})
	})

	t.Run("GetEmptyInt8SliceWithDefault", func(t *testing.T) {
		def := []int8{21, 22}
		expected := make([]int8, 0)

		t.Setenv("V1", "")

		res := env.GetInt8Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetInt8Slice(t *testing.T) {
	t.Run("MustGetAbsentInt8Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt8Slice("V1")
		})
	})

	t.Run("MustGetValidInt8Slice", func(t *testing.T) {
		expected := []int8{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetInt8Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidInt8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetInt8Slice("V1")
		})
	})

	t.Run("MustGetEmptyInt8Slice", func(t *testing.T) {
		expected := make([]int8, 0)

		t.Setenv("V1", "")

		res := env.MustGetInt8Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

func TestLookupInt8Slice(t *testing.T) {
	t.Run("LookupAbsentInt8Slice", func(t *testing.T) {
		_, err := env.LookupInt8Slice("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,3")

		res, err := env.LookupInt8Slice("V1")
		assertNoError(t, err)
		assertEqualSlices(t, []int8{1, 2, 3}, res)
	})

	t.Run("LookupInvalidInt8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		_, err := env.LookupInt8Slice("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
	})
}

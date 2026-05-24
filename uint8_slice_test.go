package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint8Slice(t *testing.T) {
	def := []uint8{21, 22}

	t.Run("GetAbsentUInt8SliceWithDefault", func(t *testing.T) {
		res := env.GetUint8Slice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidUInt8SliceWithDefault", func(t *testing.T) {
		expected := []uint8{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetUint8Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetInvalidUInt8SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.GetUint8Slice("V1", def)
		})
	})

	t.Run("GetInvalidUInt8SliceWithDefault2", func(t *testing.T) {
		expected := make([]uint8, 0)

		t.Setenv("V1", "")

		res := env.GetUint8Slice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetUint8Slice(t *testing.T) {
	t.Run("MustGetAbsentUInt8Slice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint8Slice("V1")
		})
	})

	t.Run("MustGetValidUInt8Slice", func(t *testing.T) {
		expected := []uint8{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetUint8Slice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetInvalidUInt8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assertPanics(t, func() {
			env.MustGetUint8Slice("V1")
		})
	})

	t.Run("MustGetEmptyUInt8Slice", func(t *testing.T) {
		expected := make([]uint8, 0)

		t.Setenv("V1", "")

		res := env.MustGetUint8Slice("V1")
		assertEqualSlices(t, expected, res)
	})
}

func TestLookupUint8Slice(t *testing.T) {
	t.Run("LookupAbsentUint8Slice", func(t *testing.T) {
		_, err := env.LookupUint8Slice("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidUint8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,3")

		res, err := env.LookupUint8Slice("V1")
		assertNoError(t, err)
		assertEqualSlices(t, []uint8{1, 2, 3}, res)
	})

	t.Run("LookupInvalidUint8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		_, err := env.LookupUint8Slice("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
	})
}

package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint(t *testing.T) {
	def := uint(12)

	t.Run("GetAbsentUIntWithDefault", func(t *testing.T) {
		res := env.GetUint("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidUIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetUint("V1", def)
		})
	})

	t.Run("GetValidUIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint("V1", def)
		assertEqual(t, uint(14), res)
	})
}

func TestMustGetUint(t *testing.T) {
	t.Run("MustGetAbsentUInt", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint("V1")
		})
	})

	t.Run("MustGetInvalidUInt", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint("V1")
		})
	})

	t.Run("MustGetValidUInt", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint("V1")
		assertEqual(t, uint(14), res)
	})
}

func TestLookupUint(t *testing.T) {
	t.Run("LookupAbsentUint", func(t *testing.T) {
		_, err := env.LookupUint("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidUint", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupUint("V1")
		assertNoError(t, err)
		assertEqual(t, uint(14), res)
	})

	t.Run("LookupInvalidUint", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupUint("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

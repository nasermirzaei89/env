package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt32(t *testing.T) {
	def := int32(12)

	t.Run("GetAbsentInt32WithDefault", func(t *testing.T) {
		res := env.GetInt32("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidInt32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetInt32("V1", def)
		})
	})

	t.Run("GetValidInt32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt32("V1", def)
		assertEqual(t, int32(14), res)
	})
}

func TestMustGetInt32(t *testing.T) {
	t.Run("MustGetAbsentInt32", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt32("V1")
		})
	})

	t.Run("MustGetInvalidInt32", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")

			env.MustGetInt32("V1")
		})
	})

	t.Run("MustGetValidInt32", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt32("V1")
		assertEqual(t, int32(14), res)
	})
}

func TestLookupInt32(t *testing.T) {
	t.Run("LookupAbsentInt32", func(t *testing.T) {
		_, err := env.LookupInt32("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt32", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupInt32("V1")
		assertNoError(t, err)
		assertEqual(t, int32(14), res)
	})

	t.Run("LookupInvalidInt32", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupInt32("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

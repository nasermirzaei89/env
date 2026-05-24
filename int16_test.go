package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt16(t *testing.T) {
	def := int16(12)

	t.Run("GetAbsentInt16WithDefault", func(t *testing.T) {
		res := env.GetInt16("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetInt16("V1", def)
		})
	})

	t.Run("GetValidInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt16("V1", def)
		assertEqual(t, int16(14), res)
	})
}

func TestMustGetInt16(t *testing.T) {
	t.Run("MustGetAbsentInt16", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt16("V1")
		})
	})

	t.Run("MustGetInvalidInt16", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")

			env.MustGetInt16("V1")
		})
	})

	t.Run("MustGetValidInt16", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt16("V1")
		assertEqual(t, int16(14), res)
	})
}

func TestLookupInt16(t *testing.T) {
	t.Run("LookupAbsentInt16", func(t *testing.T) {
		_, err := env.LookupInt16("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt16", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupInt16("V1")
		assertNoError(t, err)
		assertEqual(t, int16(14), res)
	})

	t.Run("LookupInvalidInt16", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupInt16("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetFloat64(t *testing.T) {
	def := 12.5

	t.Run("GetAbsentFloat64WithDefault", func(t *testing.T) {
		res := env.GetFloat64("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidFloat64WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetFloat64("V1", def)
		})
	})

	t.Run("GetValidFloat64WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14.5")

		res := env.GetFloat64("V1", def)
		assertEqual(t, 14.5, res)
	})
}

func TestMustGetFloat64(t *testing.T) {
	t.Run("MustGetAbsentFloat64", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetFloat64("V1")
		})
	})

	t.Run("MustGetInvalidFloat64", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")

			env.MustGetFloat64("V1")
		})
	})

	t.Run("MustGetValidFloat64", func(t *testing.T) {
		t.Setenv("V1", "14.5")

		res := env.MustGetFloat64("V1")
		assertEqual(t, 14.5, res)
	})
}

func TestLookupFloat64(t *testing.T) {
	t.Run("LookupAbsentFloat64", func(t *testing.T) {
		_, err := env.LookupFloat64("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidFloat64", func(t *testing.T) {
		t.Setenv("V1", "14.5")

		res, err := env.LookupFloat64("V1")
		assertNoError(t, err)
		assertEqual(t, 14.5, res)
	})

	t.Run("LookupInvalidFloat64", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupFloat64("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

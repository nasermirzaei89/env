package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt64(t *testing.T) {
	def := int64(12)

	t.Run("GetAbsentInt64WithDefault", func(t *testing.T) {
		res := env.GetInt64("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidInt64WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetInt64("V1", def)
		})
	})

	t.Run("GetValidInt64WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt64("V1", def)
		assertEqual(t, int64(14), res)
	})
}

func TestMustGetInt64(t *testing.T) {
	t.Run("MustGetAbsentInt64", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt64("V1")
		})
	})

	t.Run("MustGetInvalidInt64", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetInt64("V1")
		})
	})

	t.Run("MustGetValidInt64", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt64("V1")
		assertEqual(t, int64(14), res)
	})
}

func TestLookupInt64(t *testing.T) {
	t.Run("LookupAbsentInt64", func(t *testing.T) {
		_, err := env.LookupInt64("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt64", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupInt64("V1")
		assertNoError(t, err)
		assertEqual(t, int64(14), res)
	})

	t.Run("LookupInvalidInt64", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupInt64("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

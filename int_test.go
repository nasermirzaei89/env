package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt(t *testing.T) {
	def := 12

	t.Run("GetAbsentIntWithDefault", func(t *testing.T) {
		res := env.GetInt("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetInt("V1", def)
		})
	})

	t.Run("GetValidIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt("V1", def)
		assertEqual(t, 14, res)
	})
}

func TestMustGetInt(t *testing.T) {
	t.Run("MustGetAbsentInt", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt("V1")
		})
	})

	t.Run("MustGetInvalidInt", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetInt("V1")
		})
	})

	t.Run("MustGetValidInt", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt("V1")
		assertEqual(t, 14, res)
	})
}

func TestLookupInt(t *testing.T) {
	t.Run("LookupAbsentInt", func(t *testing.T) {
		_, err := env.LookupInt("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupInt("V1")
		assertNoError(t, err)
		assertEqual(t, 14, res)
	})

	t.Run("LookupInvalidInt", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupInt("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

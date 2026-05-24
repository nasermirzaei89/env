package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt8(t *testing.T) {
	def := int8(12)

	t.Run("GetAbsentInt8WithDefault", func(t *testing.T) {
		res := env.GetInt8("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidInt8WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetInt8("V1", def)
		})
	})

	t.Run("GetValidInt8WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt8("V1", def)
		assertEqual(t, int8(14), res)
	})
}

func TestMustGetInt8(t *testing.T) {
	t.Run("MustGetAbsentInt8", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt8("V1")
		})
	})

	t.Run("MustGetInvalidInt8", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetInt8("V1")
		})
	})

	t.Run("MustGetValidInt8", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt8("V1")
		assertEqual(t, int8(14), res)
	})
}

func TestLookupInt8(t *testing.T) {
	t.Run("LookupAbsentInt8", func(t *testing.T) {
		_, err := env.LookupInt8("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidInt8", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupInt8("V1")
		assertNoError(t, err)
		assertEqual(t, int8(14), res)
	})

	t.Run("LookupInvalidInt8", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupInt8("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint8(t *testing.T) {
	def := uint8(12)

	t.Run("GetAbsentUInt8WithDefault", func(t *testing.T) {
		res := env.GetUint8("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidUInt8WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetUint8("V1", def)
		})
	})

	t.Run("GetValidUInt8WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint8("V1", def)
		assertEqual(t, uint8(14), res)
	})
}

func TestMustGetUint8(t *testing.T) {
	t.Run("MustGetAbsentUInt8", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint8("V1")
		})
	})

	t.Run("MustGetInvalidUInt8", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint8("V1")
		})
	})

	t.Run("MustGetValidUInt8", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint8("V1")
		assertEqual(t, uint8(14), res)
	})
}

func TestLookupUint8(t *testing.T) {
	t.Run("LookupAbsentUint8", func(t *testing.T) {
		_, err := env.LookupUint8("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidUint8", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupUint8("V1")
		assertNoError(t, err)
		assertEqual(t, uint8(14), res)
	})

	t.Run("LookupInvalidUint8", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupUint8("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

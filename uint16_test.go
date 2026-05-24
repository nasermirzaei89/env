package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint16(t *testing.T) {
	def := uint16(12)

	t.Run("GetAbsentUInt16WithDefault", func(t *testing.T) {
		res := env.GetUint16("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidUInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetUint16("V1", def)
		})
	})

	t.Run("GetValidUInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint16("V1", def)
		assertEqual(t, uint16(14), res)
	})
}

func TestMustGetUint16(t *testing.T) {
	t.Run("MustGetAbsentUInt16", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint16("V1")
		})
	})

	t.Run("MustGetInvalidUInt16", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint16("V1")
		})
	})

	t.Run("MustGetValidUInt16", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint16("V1")
		assertEqual(t, uint16(14), res)
	})
}

func TestLookupUint16(t *testing.T) {
	t.Run("LookupAbsentUint16", func(t *testing.T) {
		_, err := env.LookupUint16("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidUint16", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupUint16("V1")
		assertNoError(t, err)
		assertEqual(t, uint16(14), res)
	})

	t.Run("LookupInvalidUint16", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupUint16("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint32(t *testing.T) {
	def := uint32(12)

	t.Run("GetAbsentUInt32WithDefault", func(t *testing.T) {
		res := env.GetUint32("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidUInt32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetUint32("V1", def)
		})
	})

	t.Run("GetValidUInt32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint32("V1", def)
		assertEqual(t, uint32(14), res)
	})
}

func TestMustGetUint32(t *testing.T) {
	t.Run("MustGetAbsentUInt32", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint32("V1")
		})
	})

	t.Run("MustGetInvalidUInt32", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint32("V1")
		})
	})

	t.Run("MustGetValidUInt32", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint32("V1")
		assertEqual(t, uint32(14), res)
	})
}

func TestLookupUint32(t *testing.T) {
	t.Run("LookupAbsentUint32", func(t *testing.T) {
		_, err := env.LookupUint32("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidUint32", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupUint32("V1")
		assertNoError(t, err)
		assertEqual(t, uint32(14), res)
	})

	t.Run("LookupInvalidUint32", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupUint32("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

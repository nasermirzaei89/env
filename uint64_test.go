package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint64(t *testing.T) {
	def := uint64(12)

	t.Run("GetAbsentUInt64WithDefault", func(t *testing.T) {
		res := env.GetUint64("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidUInt64WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetUint64("V1", def)
		})
	})

	t.Run("GetValidUInt64WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint64("V1", def)
		assertEqual(t, uint64(14), res)
	})
}

func TestMustGetUint64(t *testing.T) {
	t.Run("MustGetAbsentUInt64", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint64("V1")
		})
	})

	t.Run("MustGetInvalidUInt64", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint64("V1")
		})
	})

	t.Run("MustGetValidUInt64", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint64("V1")
		assertEqual(t, uint64(14), res)
	})
}

func TestLookupUint64(t *testing.T) {
	t.Run("LookupAbsentUint64", func(t *testing.T) {
		_, err := env.LookupUint64("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidUint64", func(t *testing.T) {
		t.Setenv("V1", "14")

		res, err := env.LookupUint64("V1")
		assertNoError(t, err)
		assertEqual(t, uint64(14), res)
	})

	t.Run("LookupInvalidUint64", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupUint64("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

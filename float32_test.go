package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetFloat32(t *testing.T) {
	t.Run("GetAbsentFloat32WithDefault", func(t *testing.T) {
		def := float32(12.5)

		res := env.GetFloat32("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidFloat32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetFloat32("V1", float32(12.5))
		})
	})

	t.Run("GetValidFloat32WithDefault", func(t *testing.T) {
		def := float32(12.5)

		t.Setenv("V1", "14.5")

		res := env.GetFloat32("V1", def)

		assertEqual(t, float32(14.5), res)
	})
}

func TestMustGetFloat32(t *testing.T) {
	t.Run("MustGetAbsentFloat32", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetFloat32("V1")
		})
	})

	t.Run("MustGetInvalidFloat32", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetFloat32("V1")
		})
	})

	t.Run("MustGetValidFloat32", func(t *testing.T) {
		t.Setenv("V1", "14.5")

		res := env.MustGetFloat32("V1")
		assertEqual(t, float32(14.5), res)
	})
}

func TestLookupFloat32(t *testing.T) {
	t.Run("LookupAbsentFloat32", func(t *testing.T) {
		_, err := env.LookupFloat32("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidFloat32", func(t *testing.T) {
		t.Setenv("V1", "14.5")

		res, err := env.LookupFloat32("V1")
		assertNoError(t, err)
		assertEqual(t, float32(14.5), res)
	})

	t.Run("LookupInvalidFloat32", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupFloat32("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

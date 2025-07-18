package env_test

import (
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

		def := float32(12.5)

		res := env.GetFloat32("V1", def)

		assertEqual(t, def, res)
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

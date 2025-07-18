package env_test

import (
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

		res := env.GetFloat64("V1", def)
		assertEqual(t, def, res)
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

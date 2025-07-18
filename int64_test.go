package env_test

import (
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

		res := env.GetInt64("V1", def)
		assertEqual(t, def, res)
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

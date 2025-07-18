package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetInt16(t *testing.T) {
	def := int16(12)

	t.Run("GetAbsentInt16WithDefault", func(t *testing.T) {
		res := env.GetInt16("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		res := env.GetInt16("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetValidInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt16("V1", def)
		assertEqual(t, int16(14), res)
	})
}

func TestMustGetInt16(t *testing.T) {
	t.Run("MustGetAbsentInt16", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetInt16("V1")
		})
	})

	t.Run("MustGetInvalidInt16", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")

			env.MustGetInt16("V1")
		})
	})

	t.Run("MustGetValidInt16", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt16("V1")
		assertEqual(t, int16(14), res)
	})
}

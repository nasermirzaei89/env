package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetString(t *testing.T) {
	def := "v1_default"

	t.Run("GetAbsentStringWithDefault", func(t *testing.T) {
		res := env.GetString("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetValidStringWithDefault", func(t *testing.T) {
		t.Setenv("V1", "val")

		res := env.GetString("V1", def)
		assertEqual(t, "val", res)
	})
}

func TestMustGetString(t *testing.T) {
	t.Run("MustGetAbsentString", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetString("V1")
		})
	})

	t.Run("MustGetValidString", func(t *testing.T) {
		t.Setenv("V1", "val")

		res := env.MustGetString("V1")
		assertEqual(t, "val", res)
	})
}

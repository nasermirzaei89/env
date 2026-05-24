package env_test

import (
	"errors"
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

func TestLookupString(t *testing.T) {
	t.Run("LookupAbsentString", func(t *testing.T) {
		_, err := env.LookupString("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidString", func(t *testing.T) {
		t.Setenv("V1", "val")

		res, err := env.LookupString("V1")
		assertNoError(t, err)
		assertEqual(t, "val", res)
	})
}

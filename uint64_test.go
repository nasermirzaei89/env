package env_test

import (
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

		res := env.GetUint64("V1", def)
		assertEqual(t, def, res)
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

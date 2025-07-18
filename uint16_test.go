package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetUint16(t *testing.T) {
	def := uint16(12)

	t.Run("GetAbsentUInt16WithDefault", func(t *testing.T) {
		res := env.GetUint16("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetInvalidUInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		res := env.GetUint16("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetValidUInt16WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint16("V1", def)
		assertEqual(t, uint16(14), res)
	})
}

func TestMustGetUint16(t *testing.T) {
	t.Run("MustGetAbsentUInt16", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetUint16("V1")
		})
	})

	t.Run("MustGetInvalidUInt16", func(t *testing.T) {
		assertPanics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint16("V1")
		})
	})

	t.Run("MustGetValidUInt16", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint16("V1")
		assertEqual(t, uint16(14), res)
	})
}

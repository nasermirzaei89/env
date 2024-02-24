package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat32(t *testing.T) {
	t.Run("GetAbsentFloat32WithDefault", func(t *testing.T) {
		def := float32(12.5)

		res := env.GetFloat32("V1", def)
		assert.InDelta(t, def, res, 0)
	})

	t.Run("GetInvalidFloat32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		def := float32(12.5)

		res := env.GetFloat32("V1", def)

		assert.InDelta(t, def, res, 0)
	})

	t.Run("GetValidFloat32WithDefault", func(t *testing.T) {
		def := float32(12.5)

		t.Setenv("V1", "14.5")

		res := env.GetFloat32("V1", def)

		assert.InDelta(t, float32(14.5), res, 0)
	})
}

func TestMustGetFloat32(t *testing.T) {
	t.Run("MustGetAbsentFloat32", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetFloat32("V1")
		})
	})

	t.Run("MustGetInvalidFloat32", func(t *testing.T) {
		assert.Panics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetFloat32("V1")
		})
	})

	t.Run("MustGetValidFloat32", func(t *testing.T) {
		t.Setenv("V1", "14.5")

		res := env.MustGetFloat32("V1")
		assert.InDelta(t, float32(14.5), res, 0)
	})
}

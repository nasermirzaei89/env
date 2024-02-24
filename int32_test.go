package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt32(t *testing.T) {
	def := int32(12)

	t.Run("GetAbsentInt32WithDefault", func(t *testing.T) {
		res := env.GetInt32("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetInvalidInt32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		res := env.GetInt32("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidInt32WithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt32("V1", def)
		assert.Equal(t, int32(14), res)
	})
}

func TestMustGetInt32(t *testing.T) {
	t.Run("MustGetAbsentInt32", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetInt32("V1")
		})
	})

	t.Run("MustGetInvalidInt32", func(t *testing.T) {
		assert.Panics(t, func() {
			t.Setenv("V1", "invalid")

			env.MustGetInt32("V1")
		})
	})

	t.Run("MustGetValidInt32", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt32("V1")
		assert.Equal(t, int32(14), res)
	})
}

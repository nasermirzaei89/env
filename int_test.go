package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	def := 12

	t.Run("GetAbsentIntWithDefault", func(t *testing.T) {
		res := env.GetInt("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetInvalidIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		res := env.GetInt("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetInt("V1", def)
		assert.Equal(t, 14, res)
	})
}

func TestMustGetInt(t *testing.T) {
	t.Run("MustGetAbsentInt", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetInt("V1")
		})
	})

	t.Run("MustGetInvalidInt", func(t *testing.T) {
		assert.Panics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetInt("V1")
		})
	})

	t.Run("MustGetValidInt", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetInt("V1")
		assert.Equal(t, 14, res)
	})
}

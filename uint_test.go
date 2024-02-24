package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint(t *testing.T) {
	def := uint(12)

	t.Run("GetAbsentUIntWithDefault", func(t *testing.T) {
		res := env.GetUint("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetInvalidUIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		res := env.GetUint("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidUIntWithDefault", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.GetUint("V1", def)
		assert.Equal(t, uint(14), res)
	})
}

func TestMustGetUint(t *testing.T) {
	t.Run("MustGetAbsentUInt", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetUint("V1")
		})
	})

	t.Run("MustGetInvalidUInt", func(t *testing.T) {
		assert.Panics(t, func() {
			t.Setenv("V1", "invalid")
			env.MustGetUint("V1")
		})
	})

	t.Run("MustGetValidUInt", func(t *testing.T) {
		t.Setenv("V1", "14")

		res := env.MustGetUint("V1")
		assert.Equal(t, uint(14), res)
	})
}

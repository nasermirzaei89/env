package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUintSlice(t *testing.T) {
	def := []uint{21, 22}

	t.Run("GetAbsentUIntSliceWithDefault", func(t *testing.T) {
		res := env.GetUintSlice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidUIntSliceWithDefault", func(t *testing.T) {
		expected := []uint{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetUintSlice("V1", def)
		assert.Equal(t, expected, res)
	})

	t.Run("GetInvalidUIntSliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetUintSlice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetEmptyUIntSliceWithDefault", func(t *testing.T) {
		expected := make([]uint, 0)

		t.Setenv("V1", "")

		res := env.GetUintSlice("V1", def)
		assert.Equal(t, expected, res)
	})
}

func TestMustGetUintSlice(t *testing.T) {
	t.Run("MustGetAbsentUIntSlice", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetUintSlice("V1")
		})
	})

	t.Run("MustGetValidUIntSlice", func(t *testing.T) {
		expected := []uint{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetUintSlice("V1")
		assert.Equal(t, expected, res)
	})

	t.Run("MustGetInvalidUIntSlice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assert.Panics(t, func() {
			env.MustGetUintSlice("V1")
		})
	})

	t.Run("MustGetEmptyUIntSlice", func(t *testing.T) {
		expected := make([]uint, 0)

		t.Setenv("V1", "")

		res := env.MustGetUintSlice("V1")
		assert.Equal(t, expected, res)
	})
}

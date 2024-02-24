package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt16Slice(t *testing.T) {
	def := []int16{21, 22}

	t.Run("GetAbsentInt16SliceWithDefault", func(t *testing.T) {
		res := env.GetInt16Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidInt16SliceWithDefault", func(t *testing.T) {
		expected := []int16{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetInt16Slice("V1", def)
		assert.Equal(t, expected, res)
	})

	t.Run("GetInvalidInt16SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetInt16Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetEmptyInt16SliceWithDefault", func(t *testing.T) {
		expected := make([]int16, 0)

		t.Setenv("V1", "")

		res := env.GetInt16Slice("V1", def)
		assert.Equal(t, expected, res)
	})
}

func TestMustGetInt16Slice(t *testing.T) {
	t.Run("MustGetAbsentInt16Slice", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetInt16Slice("V1")
		})
	})

	t.Run("MustGetValidInt16Slice", func(t *testing.T) {
		expected := []int16{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetInt16Slice("V1")
		assert.Equal(t, expected, res)
	})

	t.Run("MustGetInvalidInt16Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assert.Panics(t, func() {
			env.MustGetInt16Slice("V1")
		})
	})

	t.Run("MustGetEmptyInt16Slice", func(t *testing.T) {
		expected := make([]int16, 0)

		t.Setenv("V1", "")

		res := env.MustGetInt16Slice("V1")
		assert.Equal(t, expected, res)
	})
}

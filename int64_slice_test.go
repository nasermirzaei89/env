package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt64Slice(t *testing.T) {
	def := []int64{21, 22}

	t.Run("GetAbsentInt64SliceWithDefault", func(t *testing.T) {
		res := env.GetInt64Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidInt64SliceWithDefault", func(t *testing.T) {
		expected := []int64{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetInt64Slice("V1", def)
		assert.Equal(t, expected, res)
	})

	t.Run("GetInvalidInt64SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetInt64Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetEmptyInt64SliceWithDefault", func(t *testing.T) {
		expected := make([]int64, 0)

		t.Setenv("V1", "")

		res := env.GetInt64Slice("V1", def)
		assert.Equal(t, expected, res)
	})
}

func TestMustGetInt64Slice(t *testing.T) {
	t.Run("MustGetAbsentInt64Slice", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetInt64Slice("V1")
		})
	})

	t.Run("MustGetValidInt64Slice", func(t *testing.T) {
		expected := []int64{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetInt64Slice("V1")
		assert.Equal(t, expected, res)
	})

	t.Run("MustGetInvalidInt64Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assert.Panics(t, func() {
			env.MustGetInt64Slice("V1")
		})
	})

	t.Run("MustGetEmptyInt64Slice", func(t *testing.T) {
		expected := make([]int64, 0)

		t.Setenv("V1", "")

		res := env.MustGetInt64Slice("V1")
		assert.Equal(t, expected, res)
	})
}

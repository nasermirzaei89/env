package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat32Slice(t *testing.T) {
	t.Run("GetAbsentFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}

		res := env.GetFloat32Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}
		expected := []float32{31.02, 32.33, 33.33}

		t.Setenv("V1", "31.02,32.33,33.33")

		res := env.GetFloat32Slice("V1", def)
		assert.Equal(t, expected, res)
	})

	t.Run("GetInvalidFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}

		t.Setenv("V1", "1.2,2.3,Three")

		res := env.GetFloat32Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetEmptyFloat32SliceWithDefault", func(t *testing.T) {
		def := []float32{21.2, 22.3}
		expected := make([]float32, 0)

		t.Setenv("V1", "")

		res := env.GetFloat32Slice("V1", def)
		assert.Equal(t, expected, res)
	})
}

func TestMustGetFloat32Slice(t *testing.T) {
	t.Run("MustGetAbsentFloat32Slice", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetFloat32Slice("V1")
		})
	})

	t.Run("MustGetValidFloat32Slice", func(t *testing.T) {
		expected := []float32{31.02, 32.33, 33.33}

		t.Setenv("V1", "31.02,32.33,33.33")

		res := env.MustGetFloat32Slice("V1")
		assert.Equal(t, expected, res)
	})

	t.Run("MustGetInvalidFloat32Slice", func(t *testing.T) {
		t.Setenv("V1", "1.2,2.3,Three")

		assert.Panics(t, func() {
			env.MustGetFloat32Slice("V1")
		})
	})

	t.Run("MustGetEmptyFloat32Slice", func(t *testing.T) {
		expected := make([]float32, 0)

		t.Setenv("V1", "")

		res := env.MustGetFloat32Slice("V1")
		assert.Equal(t, expected, res)
	})
}

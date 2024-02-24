package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint8Slice(t *testing.T) {
	def := []uint8{21, 22}

	t.Run("GetAbsentUInt8SliceWithDefault", func(t *testing.T) {
		res := env.GetUint8Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidUInt8SliceWithDefault", func(t *testing.T) {
		expected := []uint8{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.GetUint8Slice("V1", def)
		assert.Equal(t, expected, res)
	})

	t.Run("GetInvalidUInt8SliceWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		res := env.GetUint8Slice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetInvalidUInt8SliceWithDefault2", func(t *testing.T) {
		expected := make([]uint8, 0)

		t.Setenv("V1", "")

		res := env.GetUint8Slice("V1", def)
		assert.Equal(t, expected, res)
	})
}

func TestMustGetUint8Slice(t *testing.T) {
	t.Run("MustGetAbsentUInt8Slice", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetUint8Slice("V1")
		})
	})

	t.Run("MustGetValidUInt8Slice", func(t *testing.T) {
		expected := []uint8{31, 32, 33}

		t.Setenv("V1", "31,32,33")

		res := env.MustGetUint8Slice("V1")
		assert.Equal(t, expected, res)
	})

	t.Run("MustGetInvalidUInt8Slice", func(t *testing.T) {
		t.Setenv("V1", "1,2,Three")

		assert.Panics(t, func() {
			env.MustGetUint8Slice("V1")
		})
	})

	t.Run("MustGetEmptyUInt8Slice", func(t *testing.T) {
		expected := make([]uint8, 0)

		t.Setenv("V1", "")

		res := env.MustGetUint8Slice("V1")
		assert.Equal(t, expected, res)
	})
}

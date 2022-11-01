package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint32Slice(t *testing.T) {
	def := []uint32{21, 22}

	res := env.GetUint32Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []uint32{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetUint32Slice("V1", def)
	assert.Equal(t, expected, res)

	expected = []uint32{}

	t.Setenv("V1", "")

	res = env.GetUint32Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetUint32Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint32Slice("V1")
	})

	expected := []uint32{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetUint32Slice("V1")
	assert.Equal(t, expected, res)

	expected = []uint32{}

	t.Setenv("V1", "")

	res = env.MustGetUint32Slice("V1")
	assert.Equal(t, expected, res)
}

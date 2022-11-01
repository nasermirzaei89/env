package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint8Slice(t *testing.T) {
	def := []uint8{21, 22}

	res := env.GetUint8Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []uint8{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetUint8Slice("V1", def)
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	res = env.GetUint8Slice("V1", def)
	assert.Equal(t, def, res)

	expected = []uint8{}

	t.Setenv("V1", "")

	res = env.GetUint8Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetUint8Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint8Slice("V1")
	})

	expected := []uint8{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetUint8Slice("V1")
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	assert.Panics(t, func() {
		env.MustGetUint8Slice("V1")
	})

	expected = []uint8{}

	t.Setenv("V1", "")

	res = env.MustGetUint8Slice("V1")
	assert.Equal(t, expected, res)
}

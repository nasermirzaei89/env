package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint16Slice(t *testing.T) {
	def := []uint16{21, 22}

	res := env.GetUint16Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []uint16{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetUint16Slice("V1", def)
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	res = env.GetUint16Slice("V1", def)
	assert.Equal(t, def, res)

	expected = []uint16{}

	t.Setenv("V1", "")

	res = env.GetUint16Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetUint16Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint16Slice("V1")
	})

	expected := []uint16{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetUint16Slice("V1")
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	assert.Panics(t, func() {
		env.MustGetUint16Slice("V1")
	})

	expected = []uint16{}

	t.Setenv("V1", "")

	res = env.MustGetUint16Slice("V1")
	assert.Equal(t, expected, res)
}

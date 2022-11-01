package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt8Slice(t *testing.T) {
	def := []int8{21, 22}

	res := env.GetInt8Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []int8{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetInt8Slice("V1", def)
	assert.Equal(t, expected, res)

	expected = []int8{}

	t.Setenv("V1", "")

	res = env.GetInt8Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetInt8Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt8Slice("V1")
	})

	expected := []int8{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetInt8Slice("V1")
	assert.Equal(t, expected, res)

	expected = []int8{}

	t.Setenv("V1", "")

	res = env.MustGetInt8Slice("V1")
	assert.Equal(t, expected, res)
}

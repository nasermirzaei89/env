package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt16Slice(t *testing.T) {
	def := []int16{21, 22}

	res := env.GetInt16Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []int16{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetInt16Slice("V1", def)
	assert.Equal(t, expected, res)

	expected = []int16{}

	t.Setenv("V1", "")

	res = env.GetInt16Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetInt16Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt16Slice("V1")
	})

	expected := []int16{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetInt16Slice("V1")
	assert.Equal(t, expected, res)

	expected = []int16{}

	t.Setenv("V1", "")

	res = env.MustGetInt16Slice("V1")
	assert.Equal(t, expected, res)
}

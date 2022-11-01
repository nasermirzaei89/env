package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUintSlice(t *testing.T) {
	def := []uint{21, 22}

	res := env.GetUintSlice("V1", def)
	assert.Equal(t, def, res)

	expected := []uint{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetUintSlice("V1", def)
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	res = env.GetUintSlice("V1", def)
	assert.Equal(t, def, res)

	expected = []uint{}

	t.Setenv("V1", "")

	res = env.GetUintSlice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetUintSlice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUintSlice("V1")
	})

	expected := []uint{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetUintSlice("V1")
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	assert.Panics(t, func() {
		env.MustGetUintSlice("V1")
	})

	expected = []uint{}

	t.Setenv("V1", "")

	res = env.MustGetUintSlice("V1")
	assert.Equal(t, expected, res)
}

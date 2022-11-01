package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetIntSlice(t *testing.T) {
	def := []int{21, 22}

	res := env.GetIntSlice("V1", def)
	assert.Equal(t, def, res)

	expected := []int{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetIntSlice("V1", def)
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	res = env.GetIntSlice("V1", def)
	assert.Equal(t, def, res)

	expected = []int{}

	t.Setenv("V1", "")

	res = env.GetIntSlice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetIntSlice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetIntSlice("V1")
	})

	expected := []int{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetIntSlice("V1")
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1,2,Three")

	assert.Panics(t, func() {
		env.MustGetIntSlice("V1")
	})

	expected = []int{}

	t.Setenv("V1", "")

	res = env.MustGetIntSlice("V1")
	assert.Equal(t, expected, res)
}

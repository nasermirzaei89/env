package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt32Slice(t *testing.T) {
	def := []int32{21, 22}

	res := env.GetInt32Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []int32{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetInt32Slice("V1", def)
	assert.Equal(t, expected, res)

	expected = []int32{}

	t.Setenv("V1", "")

	res = env.GetInt32Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetInt32Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt32Slice("V1")
	})

	expected := []int32{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetInt32Slice("V1")
	assert.Equal(t, expected, res)

	expected = []int32{}

	t.Setenv("V1", "")

	res = env.MustGetInt32Slice("V1")
	assert.Equal(t, expected, res)
}

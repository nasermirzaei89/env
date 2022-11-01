package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt64Slice(t *testing.T) {
	def := []int64{21, 22}

	res := env.GetInt64Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []int64{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res = env.GetInt64Slice("V1", def)
	assert.Equal(t, expected, res)

	expected = []int64{}

	t.Setenv("V1", "")

	res = env.GetInt64Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetInt64Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt64Slice("V1")
	})

	expected := []int64{31, 32, 33}

	t.Setenv("V1", "31,32,33")

	res := env.MustGetInt64Slice("V1")
	assert.Equal(t, expected, res)

	expected = []int64{}

	t.Setenv("V1", "")

	res = env.MustGetInt64Slice("V1")
	assert.Equal(t, expected, res)
}

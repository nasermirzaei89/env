package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat64Slice(t *testing.T) {
	def := []float64{21.2, 22.3}

	res := env.GetFloat64Slice("V1", def)
	assert.Equal(t, def, res)

	expected := []float64{31.02, 32.33, 33.33}

	t.Setenv("V1", "31.02,32.33,33.33")

	res = env.GetFloat64Slice("V1", def)
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1.2,2.3,Three")

	res = env.GetFloat64Slice("V1", def)
	assert.Equal(t, def, res)

	expected = []float64{}

	t.Setenv("V1", "")

	res = env.GetFloat64Slice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetFloat64Slice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetFloat64Slice("V1")
	})

	expected := []float64{31.02, 32.33, 33.33}

	t.Setenv("V1", "31.02,32.33,33.33")

	res := env.MustGetFloat64Slice("V1")
	assert.Equal(t, expected, res)

	t.Setenv("V1", "1.2,2.3,Three")

	assert.Panics(t, func() {
		env.MustGetFloat64Slice("V1")
	})

	expected = []float64{}

	t.Setenv("V1", "")

	res = env.MustGetFloat64Slice("V1")
	assert.Equal(t, expected, res)
}

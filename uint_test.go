package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint(t *testing.T) {
	def := uint(12)

	res := env.GetUint("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetUint("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetUint("V1", def)
	assert.Equal(t, uint(14), res)
}

func TestMustGetUint(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetUint("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetUint("V1")
	assert.Equal(t, uint(14), res)
}

package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat64(t *testing.T) {
	def := 12.5

	res := env.GetFloat64("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetFloat64("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14.5")

	res = env.GetFloat64("V1", def)
	assert.Equal(t, 14.5, res)
}

func TestMustGetFloat64(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetFloat64("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")

		env.MustGetFloat64("V1")
	})

	t.Setenv("V1", "14.5")

	res := env.MustGetFloat64("V1")
	assert.Equal(t, 14.5, res)
}

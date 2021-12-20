package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	def := 12

	res := env.GetInt("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetInt("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetInt("V1", def)
	assert.Equal(t, 14, res)
}

func TestMustGetInt(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetInt("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetInt("V1")
	assert.Equal(t, 14, res)
}

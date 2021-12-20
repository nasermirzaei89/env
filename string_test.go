package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	def := "v1_default"

	res := env.GetString("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "val")

	res = env.GetString("V1", def)
	assert.Equal(t, "val", res)
}

func TestMustGetString(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetString("V1")
	})

	t.Setenv("V1", "val")

	res := env.MustGetString("V1")
	assert.Equal(t, "val", res)
}

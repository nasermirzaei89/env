package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt8(t *testing.T) {
	def := int8(12)

	res := env.GetInt8("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetInt8("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetInt8("V1", def)
	assert.Equal(t, int8(14), res)
}

func TestMustGetInt8(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt8("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetInt8("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetInt8("V1")
	assert.Equal(t, int8(14), res)
}

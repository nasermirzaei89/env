package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint8(t *testing.T) {
	def := uint8(12)

	res := env.GetUint8("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetUint8("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetUint8("V1", def)
	assert.Equal(t, uint8(14), res)
}

func TestMustGetUint8(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint8("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetUint8("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetUint8("V1")
	assert.Equal(t, uint8(14), res)
}

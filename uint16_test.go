package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint16(t *testing.T) {
	def := uint16(12)

	res := env.GetUint16("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetUint16("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetUint16("V1", def)
	assert.Equal(t, uint16(14), res)
}

func TestMustGetUint16(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint16("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetUint16("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetUint16("V1")
	assert.Equal(t, uint16(14), res)
}

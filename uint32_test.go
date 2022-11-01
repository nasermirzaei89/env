package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint32(t *testing.T) {
	def := uint32(12)

	res := env.GetUint32("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetUint32("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetUint32("V1", def)
	assert.Equal(t, uint32(14), res)
}

func TestMustGetUint32(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint32("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetUint32("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetUint32("V1")
	assert.Equal(t, uint32(14), res)
}

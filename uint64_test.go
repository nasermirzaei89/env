package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetUint64(t *testing.T) {
	def := uint64(12)

	res := env.GetUint64("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetUint64("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetUint64("V1", def)
	assert.Equal(t, uint64(14), res)
}

func TestMustGetUint64(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetUint64("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetUint64("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetUint64("V1")
	assert.Equal(t, uint64(14), res)
}

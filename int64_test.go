package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetInt64(t *testing.T) {
	def := int64(12)

	res := env.GetInt64("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "invalid")

	res = env.GetInt64("V1", def)
	assert.Equal(t, def, res)

	t.Setenv("V1", "14")

	res = env.GetInt64("V1", def)
	assert.Equal(t, int64(14), res)
}

func TestMustGetInt64(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetInt64("V1")
	})

	assert.Panics(t, func() {
		t.Setenv("V1", "invalid")
		env.MustGetInt64("V1")
	})

	t.Setenv("V1", "14")

	res := env.MustGetInt64("V1")
	assert.Equal(t, int64(14), res)
}

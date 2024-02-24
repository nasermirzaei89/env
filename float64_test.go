package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat64(t *testing.T) {
	def := 12.5

	res := env.GetFloat64("V1", def)
	assert.InDelta(t, def, res, 0)

	t.Setenv("V1", "invalid")

	res = env.GetFloat64("V1", def)
	assert.InDelta(t, def, res, 0)

	t.Setenv("V1", "14.5")

	res = env.GetFloat64("V1", def)
	assert.InDelta(t, 14.5, res, 0)
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
	assert.InDelta(t, 14.5, res, 0)
}

package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetFloat32(t *testing.T) {
	os.Clearenv()

	def := float32(12.5)
	res := env.GetFloat32("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "invalid")
	res = env.GetFloat32("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "14.5")
	res = env.GetFloat32("V1", def)
	assert.Equal(t, float32(14.5), res)
}

func TestMustGetFloat32(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetFloat32("V1")
	})

	assert.Panics(t, func() {
		_ = os.Setenv("V1", "invalid")
		env.MustGetFloat32("V1")
	})

	_ = os.Setenv("V1", "14.5")
	res := env.MustGetFloat32("V1")
	assert.Equal(t, float32(14.5), res)
}

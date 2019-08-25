package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetInt(t *testing.T) {
	os.Clearenv()

	def := 12
	res := env.GetInt("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "invalid")
	res = env.GetInt("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "14")
	res = env.GetInt("V1", def)
	assert.Equal(t, 14, res)
}

func TestMustGetInt(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetInt("V1")
	})

	assert.Panics(t, func() {
		_ = os.Setenv("V1", "invalid")
		env.MustGetInt("V1")
	})

	_ = os.Setenv("V1", "14")
	res := env.MustGetInt("V1")
	assert.Equal(t, 14, res)
}

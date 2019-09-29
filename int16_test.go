package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetInt16(t *testing.T) {
	os.Clearenv()

	def := int16(12)
	res := env.GetInt16("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "invalid")
	res = env.GetInt16("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "14")
	res = env.GetInt16("V1", def)
	assert.Equal(t, int16(14), res)
}

func TestMustGetInt16(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetInt16("V1")
	})

	assert.Panics(t, func() {
		_ = os.Setenv("V1", "invalid")
		env.MustGetInt16("V1")
	})

	_ = os.Setenv("V1", "14")
	res := env.MustGetInt16("V1")
	assert.Equal(t, int16(14), res)
}

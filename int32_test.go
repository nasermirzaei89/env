package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetInt32(t *testing.T) {
	os.Clearenv()

	def := int32(12)
	res := env.GetInt32("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "invalid")
	res = env.GetInt32("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "14")
	res = env.GetInt32("V1", def)
	assert.Equal(t, int32(14), res)
}

func TestMustGetInt32(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetInt32("V1")
	})

	assert.Panics(t, func() {
		_ = os.Setenv("V1", "invalid")
		env.MustGetInt32("V1")
	})

	_ = os.Setenv("V1", "14")
	res := env.MustGetInt32("V1")
	assert.Equal(t, int32(14), res)
}

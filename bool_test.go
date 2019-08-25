package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetBool(t *testing.T) {
	os.Clearenv()

	res := env.GetBool("V1", true)
	assert.True(t, res)

	_ = os.Setenv("V1", "false")
	res = env.GetBool("V1", true)
	assert.False(t, res)

	_ = os.Setenv("V1", "0")
	res = env.GetBool("V1", true)
	assert.False(t, res)

	_ = os.Setenv("V1", "true")
	res = env.GetBool("V1", false)
	assert.True(t, res)

	_ = os.Setenv("V1", "")
	res = env.GetBool("V1", false)
	assert.True(t, res)

	_ = os.Setenv("V1", "1")
	res = env.GetBool("V1", false)
	assert.True(t, res)
}

func TestMustGetBool(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetBool("V1")
	})

	_ = os.Setenv("V1", "false")
	res := env.MustGetBool("V1")
	assert.False(t, res)

	_ = os.Setenv("V1", "0")
	res = env.MustGetBool("V1")
	assert.False(t, res)

	_ = os.Setenv("V1", "true")
	res = env.MustGetBool("V1")
	assert.True(t, res)

	_ = os.Setenv("V1", "1")
	res = env.MustGetBool("V1")
	assert.True(t, res)
}

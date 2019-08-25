package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Clearenv()

	def := "v1_default"
	res := env.GetString("V1", def)
	assert.Equal(t, def, res)

	_ = os.Setenv("V1", "val")
	res = env.GetString("V1", def)
	assert.Equal(t, "val", res)
}

func TestMustGetString(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetString("V1")
	})

	_ = os.Setenv("V1", "val")
	res := env.MustGetString("V1")
	assert.Equal(t, "val", res)
}

package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGetStringSlice(t *testing.T) {
	os.Clearenv()

	def := []string{"foo", "bar"}
	res := env.GetStringSlice("V1", def)
	assert.Equal(t, def, res)

	expected := []string{"foo", "bar", "baz"}
	_ = os.Setenv("V1", strings.Join(expected, ","))
	res = env.GetStringSlice("V1", def)
	assert.Equal(t, expected, res)

	expected = []string{}
	_ = os.Setenv("V1", strings.Join(expected, ","))
	res = env.GetStringSlice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetStringSlice(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		env.MustGetStringSlice("V1")
	})

	expected := []string{"foo", "bar", "baz"}
	_ = os.Setenv("V1", strings.Join(expected, ","))
	res := env.MustGetStringSlice("V1")
	assert.Equal(t, expected, res)

	expected = []string{}
	_ = os.Setenv("V1", strings.Join(expected, ","))
	res = env.MustGetStringSlice("V1")
	assert.Equal(t, expected, res)
}

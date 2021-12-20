package env_test

import (
	"strings"
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetStringSlice(t *testing.T) {
	def := []string{"foo", "bar"}

	res := env.GetStringSlice("V1", def)
	assert.Equal(t, def, res)

	expected := []string{"foo", "bar", "baz"}
	t.Setenv("V1", strings.Join(expected, ","))

	res = env.GetStringSlice("V1", def)
	assert.Equal(t, expected, res)

	expected = []string{}
	t.Setenv("V1", strings.Join(expected, ","))

	res = env.GetStringSlice("V1", def)
	assert.Equal(t, expected, res)
}

func TestMustGetStringSlice(t *testing.T) {
	assert.Panics(t, func() {
		env.MustGetStringSlice("V1")
	})

	expected := []string{"foo", "bar", "baz"}
	t.Setenv("V1", strings.Join(expected, ","))

	res := env.MustGetStringSlice("V1")
	assert.Equal(t, expected, res)

	expected = []string{}
	t.Setenv("V1", strings.Join(expected, ","))

	res = env.MustGetStringSlice("V1")
	assert.Equal(t, expected, res)
}

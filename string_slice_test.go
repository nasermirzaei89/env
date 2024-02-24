package env_test

import (
	"strings"
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetStringSlice(t *testing.T) {
	def := []string{"foo", "bar"}

	t.Run("GetAbsentStringSliceWithDefault", func(t *testing.T) {
		res := env.GetStringSlice("V1", def)
		assert.Equal(t, def, res)
	})

	t.Run("GetValidStringSliceWithDefault", func(t *testing.T) {
		expected := []string{"foo", "bar", "baz"}
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.GetStringSlice("V1", def)
		assert.Equal(t, expected, res)
	})

	t.Run("GetEmptyStringSliceWithDefault", func(t *testing.T) {
		expected := make([]string, 0)
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.GetStringSlice("V1", def)
		assert.Equal(t, expected, res)
	})
}

func TestMustGetStringSlice(t *testing.T) {
	t.Run("MustGetAbsentStringSlice", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetStringSlice("V1")
		})
	})

	t.Run("MustGetValidStringSlice", func(t *testing.T) {
		expected := []string{"foo", "bar", "baz"}
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.MustGetStringSlice("V1")
		assert.Equal(t, expected, res)
	})

	t.Run("MustGetEmptyStringSlice", func(t *testing.T) {
		expected := make([]string, 0)
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.MustGetStringSlice("V1")
		assert.Equal(t, expected, res)
	})
}

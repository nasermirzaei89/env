package env_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestGetStringSlice(t *testing.T) {
	def := []string{"foo", "bar"}

	t.Run("GetAbsentStringSliceWithDefault", func(t *testing.T) {
		res := env.GetStringSlice("V1", def)
		assertEqualSlices(t, def, res)
	})

	t.Run("GetValidStringSliceWithDefault", func(t *testing.T) {
		expected := []string{"foo", "bar", "baz"}
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.GetStringSlice("V1", def)
		assertEqualSlices(t, expected, res)
	})

	t.Run("GetEmptyStringSliceWithDefault", func(t *testing.T) {
		expected := make([]string, 0)
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.GetStringSlice("V1", def)
		assertEqualSlices(t, expected, res)
	})
}

func TestMustGetStringSlice(t *testing.T) {
	t.Run("MustGetAbsentStringSlice", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetStringSlice("V1")
		})
	})

	t.Run("MustGetValidStringSlice", func(t *testing.T) {
		expected := []string{"foo", "bar", "baz"}
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.MustGetStringSlice("V1")
		assertEqualSlices(t, expected, res)
	})

	t.Run("MustGetEmptyStringSlice", func(t *testing.T) {
		expected := make([]string, 0)
		t.Setenv("V1", strings.Join(expected, ","))

		res := env.MustGetStringSlice("V1")
		assertEqualSlices(t, expected, res)
	})
}

func TestLookupStringSlice(t *testing.T) {
	t.Run("LookupAbsentStringSlice", func(t *testing.T) {
		_, err := env.LookupStringSlice("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidStringSlice", func(t *testing.T) {
		expected := []string{"foo", "bar", "baz"}
		t.Setenv("V1", strings.Join(expected, ","))

		res, err := env.LookupStringSlice("V1")
		assertNoError(t, err)
		assertEqualSlices(t, expected, res)
	})
}

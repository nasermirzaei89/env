package env_test

import (
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	t.Run("GetAbsentBoolWithDefault", func(t *testing.T) {
		res := env.GetBool("V1", true)
		assert.True(t, res)
	})

	t.Run("GetPresentBoolWithDefault", func(t *testing.T) {
		t.Setenv("V1", "false")

		res := env.GetBool("V1", true)
		assert.False(t, res)
	})

	t.Run("GetZeroAsBoolWithDefault", func(t *testing.T) {
		t.Setenv("V1", "0")

		res := env.GetBool("V1", true)
		assert.False(t, res)
	})

	t.Run("GetTrueStringAsBoolWithDefault", func(t *testing.T) {
		t.Setenv("V1", "true")

		res := env.GetBool("V1", false)
		assert.True(t, res)
	})

	t.Run("GetEmptyAsBoolWithDefault", func(t *testing.T) {
		t.Setenv("V1", "")

		res := env.GetBool("V1", false)
		assert.True(t, res)
	})

	t.Run("GetOneAsBoolWithDefault", func(t *testing.T) {
		t.Setenv("V1", "1")

		res := env.GetBool("V1", false)
		assert.True(t, res)
	})
}

func TestMustGetBool(t *testing.T) {
	t.Run("MustGetAbsentBool", func(t *testing.T) {
		assert.Panics(t, func() {
			env.MustGetBool("V1")
		})
	})

	t.Run("MustGetFalseStringAsBool", func(t *testing.T) {
		t.Setenv("V1", "false")

		res := env.MustGetBool("V1")
		assert.False(t, res)
	})

	t.Run("MustGetZeroAsBool", func(t *testing.T) {
		t.Setenv("V1", "0")

		res := env.MustGetBool("V1")
		assert.False(t, res)
	})

	t.Run("MustGetTrueStringAsBool", func(t *testing.T) {
		t.Setenv("V1", "true")

		res := env.MustGetBool("V1")
		assert.True(t, res)
	})

	t.Run("MustGetOneAsBool", func(t *testing.T) {
		t.Setenv("V1", "1")

		res := env.MustGetBool("V1")
		assert.True(t, res)
	})
}

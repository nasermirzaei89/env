package env_test

import (
	"testing"
	"time"

	"github.com/nasermirzaei89/env"
)

func TestGetDuration(t *testing.T) {
	def := time.Second

	t.Run("GetAbsentDurationWithDefault", func(t *testing.T) {
		res := env.GetDuration("V1", def)
		assertEqual(t, def, res)
	})

	t.Run("GetValidDurationWithDefault", func(t *testing.T) {
		t.Setenv("V1", "2s")

		res := env.GetDuration("V1", def)
		assertEqual(t, 2*time.Second, res)
	})
}

func TestMustGetDuration(t *testing.T) {
	t.Run("MustGetAbsentDuration", func(t *testing.T) {
		assertPanics(t, func() {
			env.MustGetDuration("V1")
		})
	})

	t.Run("MustGetValidDuration", func(t *testing.T) {
		t.Setenv("V1", "2s")

		res := env.MustGetDuration("V1")
		assertEqual(t, 2*time.Second, res)
	})
}

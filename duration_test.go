package env_test

import (
	"errors"
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

	t.Run("GetInvalidDurationWithDefault", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		assertPanics(t, func() {
			env.GetDuration("V1", def)
		})
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

func TestLookupDuration(t *testing.T) {
	t.Run("LookupAbsentDuration", func(t *testing.T) {
		_, err := env.LookupDuration("V1")

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "V1", notSetErr.Key)
	})

	t.Run("LookupValidDuration", func(t *testing.T) {
		t.Setenv("V1", "2s")

		res, err := env.LookupDuration("V1")
		assertNoError(t, err)
		assertEqual(t, 2*time.Second, res)
	})

	t.Run("LookupInvalidDuration", func(t *testing.T) {
		t.Setenv("V1", "invalid")

		_, err := env.LookupDuration("V1")

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "V1", invalidValueErr.Key)
		assertEqual(t, "invalid", invalidValueErr.Value)
	})
}

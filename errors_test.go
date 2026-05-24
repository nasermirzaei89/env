package env_test

import (
	"errors"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestNotSetError(t *testing.T) {
	t.Run("ErrorMessage", func(t *testing.T) {
		notSetErr := env.NotSetError{Key: "MY_VAR"}
		assertEqual(t, `environment variable "MY_VAR" not set`, notSetErr.Error())
	})

	t.Run("ErrorsAs", func(t *testing.T) {
		err := error(env.NotSetError{Key: "MY_VAR"})

		var notSetErr env.NotSetError
		assertTrue(t, errors.As(err, &notSetErr))
		assertEqual(t, "MY_VAR", notSetErr.Key)
	})
}

func TestInvalidValueError(t *testing.T) {
	t.Run("ErrorMessage", func(t *testing.T) {
		inner := errors.New("some parse error") //nolint:err113
		invalidValueErr := env.InvalidValueError{Key: "MY_VAR", Value: "bad", Err: inner}
		assertEqual(t, `environment variable "MY_VAR" has an invalid value: "bad"`, invalidValueErr.Error())
	})

	t.Run("Unwrap", func(t *testing.T) {
		inner := errors.New("some parse error") //nolint:err113
		invalidValueErr := env.InvalidValueError{Key: "MY_VAR", Value: "bad", Err: inner}
		assertTrue(t, errors.Is(invalidValueErr, inner))
	})

	t.Run("ErrorsAs", func(t *testing.T) {
		inner := errors.New("some parse error") //nolint:err113
		err := error(env.InvalidValueError{Key: "MY_VAR", Value: "bad", Err: inner})

		var invalidValueErr env.InvalidValueError
		assertTrue(t, errors.As(err, &invalidValueErr))
		assertEqual(t, "MY_VAR", invalidValueErr.Key)
		assertEqual(t, "bad", invalidValueErr.Value)
	})
}

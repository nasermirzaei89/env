package env_test

import (
	"runtime/debug"
	"slices"
	"testing"
)

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected == actual {
		return
	}

	t.Errorf("expected: %#v\n actual: %#v", expected, actual)
}

func assertTrue(t *testing.T, value bool) {
	t.Helper()

	assertEqual(t, true, value)
}

func assertFalse(t *testing.T, value bool) {
	t.Helper()

	if !value {
		return
	}

	assertEqual(t, false, value)
}

func assertEqualSlices[T comparable](t *testing.T, expected, actual []T) {
	t.Helper()

	if !slices.Equal(expected, actual) {
		t.Errorf("expected: %#v\n actual: %#v", expected, actual)
	}
}

func assertPanics(t *testing.T, f func()) {
	t.Helper()

	if funcDidPanic, _, _ := didPanic(f); funcDidPanic {
		return
	}

	t.Error("expected panic")
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		return
	}

	t.Errorf("unexpected error: %v", err)
}

func didPanic(f func()) (funcDidPanic bool, message any, stack string) {
	funcDidPanic = true

	defer func() {
		message = recover()

		if funcDidPanic {
			stack = string(debug.Stack())
		}
	}()

	f()

	funcDidPanic = false

	return
}

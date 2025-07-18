package env_test

import (
	"runtime/debug"
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

	var equal bool

	if len(actual) == len(expected) {
		equal = true
		for i := range actual {
			if actual[i] != expected[i] {
				equal = false
				break
			}
		}
	}

	if !equal {
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

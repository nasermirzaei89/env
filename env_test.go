package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/nasermirzaei89/env"
)

func TestEnv_String(t *testing.T) {
	v := "testing"

	t.Setenv("ENV", v)

	res := env.Environment()
	assertEqual(t, v, string(res))
}

func ExampleEnvironment() {
	_ = os.Setenv("ENV", "testing")

	fmt.Println(env.Environment())
	// Output: testing
}

func TestEnvironment(t *testing.T) {
	assertEqual(t, "", env.Environment())

	t.Setenv("ENV", "testing")

	assertEqual(t, env.Testing, env.Environment())
}

func TestIs(t *testing.T) {
	t.Run("Single Env", func(t *testing.T) {
		t.Setenv("ENV", "testing")

		assertTrue(t, env.Is(env.Testing))
	})

	t.Run("Multiple Envs", func(t *testing.T) {
		t.Setenv("ENV", "testing")

		assertTrue(t, env.Is(env.Development, env.Testing))
	})

	t.Run("Multiple Envs", func(t *testing.T) {
		t.Setenv("ENV", "testing")

		assertTrue(t, env.Is(env.Testing, env.Development))
	})
}

func ExampleIs() {
	_ = os.Setenv("ENV", "testing")

	fmt.Println(env.Is(env.Testing))
	// Output: true
}

func TestIsDevelopment(t *testing.T) {
	assertFalse(t, env.IsDevelopment())

	t.Setenv("ENV", "development")

	assertTrue(t, env.IsDevelopment())
}

func TestIsTesting(t *testing.T) {
	assertFalse(t, env.IsTesting())

	t.Setenv("ENV", "testing")

	assertTrue(t, env.IsTesting())
}

func TestIsStaging(t *testing.T) {
	assertFalse(t, env.IsStaging())

	t.Setenv("ENV", "staging")

	assertTrue(t, env.IsStaging())
}

func TestIsProduction(t *testing.T) {
	assertFalse(t, env.IsProduction())

	t.Setenv("ENV", "production")

	assertTrue(t, env.IsProduction())
}

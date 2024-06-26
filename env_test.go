package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
)

func TestEnv_String(t *testing.T) {
	v := "testing"

	t.Setenv("ENV", v)

	res := env.Environment()
	assert.EqualValues(t, v, res)
}

func ExampleEnvironment() {
	_ = os.Setenv("ENV", "testing")

	fmt.Println(env.Environment())
	// Output: testing
}

func TestEnvironment(t *testing.T) {
	assert.Zero(t, env.Environment())

	t.Setenv("ENV", "testing")

	assert.Equal(t, env.Testing, env.Environment())
}

func TestIs(t *testing.T) {
	t.Run("Single Env", func(t *testing.T) {
		t.Setenv("ENV", "testing")

		assert.True(t, env.Is(env.Testing))
	})

	t.Run("Multiple Envs", func(t *testing.T) {
		t.Setenv("ENV", "testing")

		assert.True(t, env.Is(env.Development, env.Testing))
	})

	t.Run("Multiple Envs", func(t *testing.T) {
		t.Setenv("ENV", "testing")

		assert.True(t, env.Is(env.Testing, env.Development))
	})
}

func ExampleIs() {
	_ = os.Setenv("ENV", "testing")

	fmt.Println(env.Is(env.Testing))
	// Output: true
}

func TestIsDevelopment(t *testing.T) {
	assert.False(t, env.IsDevelopment())

	t.Setenv("ENV", "development")

	assert.True(t, env.IsDevelopment())
}

func TestIsTesting(t *testing.T) {
	assert.False(t, env.IsTesting())

	t.Setenv("ENV", "testing")

	assert.True(t, env.IsTesting())
}

func TestIsStaging(t *testing.T) {
	assert.False(t, env.IsStaging())

	t.Setenv("ENV", "staging")

	assert.True(t, env.IsStaging())
}

func TestIsProduction(t *testing.T) {
	assert.False(t, env.IsProduction())

	t.Setenv("ENV", "production")

	assert.True(t, env.IsProduction())
}

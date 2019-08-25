package env_test

import (
	"github.com/nasermirzaei89/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvironment(t *testing.T) {
	os.Clearenv()

	assert.Zero(t, env.Environment())

	_ = os.Setenv("ENV", "testing")
	assert.Equal(t, env.Testing, env.Environment())
}

func TestIs(t *testing.T) {
	os.Clearenv()

	_ = os.Setenv("ENV", "testing")
	assert.True(t, env.Is(env.Testing))
}

func TestIsDevelopment(t *testing.T) {
	os.Clearenv()

	assert.False(t, env.IsDevelopment())
	_ = os.Setenv("ENV", "development")
	assert.True(t, env.IsDevelopment())
}

func TestIsTesting(t *testing.T) {
	os.Clearenv()

	assert.False(t, env.IsTesting())
	_ = os.Setenv("ENV", "testing")
	assert.True(t, env.IsTesting())
}

func TestIsStaging(t *testing.T) {
	os.Clearenv()

	assert.False(t, env.IsStaging())
	_ = os.Setenv("ENV", "staging")
	assert.True(t, env.IsStaging())
}

func TestIsProduction(t *testing.T) {
	os.Clearenv()

	assert.False(t, env.IsProduction())
	_ = os.Setenv("ENV", "production")
	assert.True(t, env.IsProduction())
}

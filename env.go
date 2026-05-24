// Package env provides functions to work with environment variables.
package env

import (
	"os"
	"slices"
)

// Env type.
type Env string

// Popular environments.
const (
	Development Env = "development"
	Testing     Env = "testing"
	Staging     Env = "staging"
	Production  Env = "production"
)

// Environment returns ENV value in environment variables.
func Environment() Env {
	return Env(os.Getenv("ENV"))
}

// Is checks whether ENV is what you expected or not.
func Is(e1 Env, ee ...Env) bool {
	ce := Environment()

	if e1 == ce {
		return true
	}

	return slices.Contains(ee, ce)
}

// IsDevelopment checks whether ENV is development or not.
func IsDevelopment() bool {
	return Is(Development)
}

// IsTesting checks whether ENV is testing or not.
func IsTesting() bool {
	return Is(Testing)
}

// IsStaging checks whether ENV is staging or not.
func IsStaging() bool {
	return Is(Staging)
}

// IsProduction checks whether ENV is production or not.
func IsProduction() bool {
	return Is(Production)
}

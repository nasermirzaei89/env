package env_test

import (
	"fmt"
	"github.com/nasermirzaei89/env"
	"os"
	"testing"
)

func except(t *testing.T, actual, excepted interface{}) {
	if fmt.Sprintf("%+v", excepted) != fmt.Sprintf("%+v", actual) {
		t.Errorf("\nexcepted: %+v\nactual:   %+v", excepted, actual)
	}
}

func notExcept(t *testing.T, actual, notExcepted interface{}) {
	if fmt.Sprintf("%+v", notExcepted) == fmt.Sprintf("%+v", actual) {
		t.Errorf("\nnot excepted: %+v\nactual:       %+v", notExcepted, actual)
	}
}

func TestEnvironment(t *testing.T) {
	os.Clearenv()
	except(t, env.Environment(), "")

	_ = os.Setenv("ENV", "testing")
	except(t, env.Environment(), env.Testing)
}

func TestIs(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("ENV", "testing")
	except(t, env.Is(env.Testing), true)
}

func TestIsDevelopment(t *testing.T) {
	os.Clearenv()
	except(t, env.IsDevelopment(), false)
	_ = os.Setenv("ENV", "development")
	except(t, env.IsDevelopment(), true)
}

func TestIsTesting(t *testing.T) {
	os.Clearenv()
	except(t, env.IsTesting(), false)
	_ = os.Setenv("ENV", "testing")
	except(t, env.IsTesting(), true)
}

func TestIsStaging(t *testing.T) {
	os.Clearenv()
	except(t, env.IsStaging(), false)
	_ = os.Setenv("ENV", "staging")
	except(t, env.IsStaging(), true)
}

func TestIsProduction(t *testing.T) {
	os.Clearenv()
	except(t, env.IsProduction(), false)
	_ = os.Setenv("ENV", "production")
	except(t, env.IsProduction(), true)
}

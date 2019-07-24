package env_test

import (
	"github.com/nasermirzaei89/env"
	"os"
	"testing"
)

func TestGetInt(t *testing.T) {
	os.Clearenv()
	def := 12
	res := env.GetInt("V1", def)
	except(t, res, def)

	_ = os.Setenv("V1", "14")
	res = env.GetInt("V1", def)
	except(t, res, "14")
}

func TestMustGetInt(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("V1", "14")
	res := env.MustGetInt("V1")
	except(t, res, 14)
}

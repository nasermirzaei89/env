package env_test

import (
	"github.com/nasermirzaei89/env"
	"os"
	"testing"
)

func TestGetFloat(t *testing.T) {
	os.Clearenv()
	def := 12.5
	res := env.GetFloat("V1", def)
	except(t, res, def)

	_ = os.Setenv("V1", "14.5")
	res = env.GetFloat("V1", def)
	except(t, res, 14.5)
}

func TestMustGetFloat(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("V1", "14.5")
	res := env.MustGetFloat("V1")
	except(t, res, 14.5)
}

package env_test

import (
	"github.com/nasermirzaei89/env"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Clearenv()
	def := "v1_default"
	res := env.GetString("V1", def)
	except(t, res, def)

	_ = os.Setenv("V1", "val")
	res = env.GetString("V1", def)
	except(t, res, "val")
}

func TestMustGetString(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("V1", "val")
	res := env.MustGetString("V1")
	except(t, res, "val")
}

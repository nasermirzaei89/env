package env_test

import (
	"github.com/nasermirzaei89/env"
	"os"
	"testing"
)

func TestGetBool(t *testing.T) {
	os.Clearenv()
	res := env.GetBool("V1", true)
	except(t, res, true)

	_ = os.Setenv("V1", "false")
	res = env.GetBool("V1", true)
	except(t, res, false)

	_ = os.Setenv("V1", "0")
	res = env.GetBool("V1", true)
	except(t, res, false)

	_ = os.Setenv("V1", "true")
	res = env.GetBool("V1", false)
	except(t, res, true)

	_ = os.Setenv("V1", "")
	res = env.GetBool("V1", false)
	except(t, res, true)

	_ = os.Setenv("V1", "1")
	res = env.GetBool("V1", false)
	except(t, res, true)
}

func TestMustGetBool(t *testing.T) {
	os.Clearenv()

	_ = os.Setenv("V1", "false")
	res := env.MustGetBool("V1")
	except(t, res, false)

	_ = os.Setenv("V1", "0")
	res = env.MustGetBool("V1")
	except(t, res, false)

	_ = os.Setenv("V1", "true")
	res = env.MustGetBool("V1")
	except(t, res, true)

	_ = os.Setenv("V1", "1")
	res = env.MustGetBool("V1")
	except(t, res, true)
}

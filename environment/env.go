package environment

import (
	"os"
)

func EnvVariable(id string, pass string) (string, string) {

	// set env variable using os package
	os.Setenv(id, "atharv")
	os.Setenv(pass, "ath123")

	// return the env variable using os package
	return os.Getenv(id), os.Getenv(pass)
}

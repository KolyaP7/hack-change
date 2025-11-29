package getenv

import (
	"os"
)

func GetValue(name string, default_val string) string {
	val := os.Getenv(name)
	if val == "" {
		val = default_val
	}
	return val

}

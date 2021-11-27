package utils

import (
	"os"
)

func GetEnv(envVariableName string, defaultValue string) string {
	var value string = os.Getenv(envVariableName)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}

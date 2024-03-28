package util

import (
	"os"
	"reflect"
)

func CheckRequiredEnvVars(env interface{}) any {
	// pass address of env struct to reflect.ValueOf
	v := reflect.ValueOf(env).Elem()

	for i := 0; i < v.NumField(); i++ {
		envVar := v.Type().Field(i).Tag.Get("env")
		envVarValue := os.Getenv(envVar)

		if envVarValue == "" {
			panic("Missing required environment variable: " + envVar)
		}

		// set the value from the environment variable to the struct field
		v.Field(i).SetString(envVarValue)
	}

	return env
}

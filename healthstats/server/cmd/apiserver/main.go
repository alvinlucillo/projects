package main

import (
	"healthstats/pkg/router"
	"healthstats/pkg/services"
	"net/http"
	"os"
	"reflect"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	env := checkRequiredEnvVars()

	l := logger.With().Str("package", "main").Str("function", "main").Logger()
	service, err := services.NewService(services.ServiceConfig{
		AWSRegion:            env.AWSRegion,
		AWSS3RoleArn:         env.AWSS3RoleArn,
		AWSAccessKey:         env.AWSAccessKey,
		AWSSecretAccessKey:   env.AWSSecretAccessKey,
		Logger:               logger,
		AWSS3FilesBucketName: env.AWSS3FilesBucketName,
	})

	if err != nil {
		l.Err(err).Msg("Error creating service")
		os.Exit(1)
	}

	r := router.NewRouter(service)
	http.ListenAndServe(":9000", r)
}

type Env struct {
	AWSRegion            string `env:"AWS_REGION"`
	AWSS3RoleArn         string `env:"AWS_S3_ROLE_ARN"`
	AWSAccessKey         string `env:"AWS_ACCESS_KEY"`
	AWSSecretAccessKey   string `env:"AWS_SECRET_ACCESS_KEY"`
	AWSS3FilesBucketName string `env:"AWS_S3_FILES_BUCKET_NAME"`
}

func checkRequiredEnvVars() Env {
	env := Env{}

	// pass address of env struct to reflect.ValueOf
	v := reflect.ValueOf(&env).Elem()

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

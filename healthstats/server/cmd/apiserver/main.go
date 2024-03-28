package main

import (
	"healthstats/pkg/router"
	"healthstats/pkg/service"
	"healthstats/pkg/util"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	env := &Env{}

	env = util.CheckRequiredEnvVars(env).(*Env)

	l := logger.With().Str("package", "main").Str("function", "main").Logger()
	svc, err := service.NewService(service.ServiceConfig{
		AWSRegion:            env.AWSRegion,
		AWSS3RoleArn:         env.AWSS3RoleArn,
		AWSAccessKey:         env.AWSAccessKey,
		AWSSecretAccessKey:   env.AWSSecretAccessKey,
		Logger:               logger,
		AWSS3FilesBucketName: env.AWSS3FilesBucketName,
		DBHost:               env.DBHost,
		DBPort:               env.DBPort,
		DBUser:               env.DBUser,
		DBPassword:           env.DBPassword,
		DBName:               env.DBName,
	})

	if err != nil {
		l.Err(err).Msg("Error creating service")
		os.Exit(1)
	}

	r := router.NewRouter(svc)
	http.ListenAndServe(":9000", r)
}

type Env struct {
	AWSRegion            string `env:"AWS_REGION"`
	AWSS3RoleArn         string `env:"AWS_S3_ROLE_ARN"`
	AWSAccessKey         string `env:"AWS_ACCESS_KEY"`
	AWSSecretAccessKey   string `env:"AWS_SECRET_ACCESS_KEY"`
	AWSS3FilesBucketName string `env:"AWS_S3_FILES_BUCKET_NAME"`
	DBHost               string `env:"DB_HOST"`
	DBPort               string `env:"DB_PORT"`
	DBUser               string `env:"DB_USER"`
	DBPassword           string `env:"DB_PASSWORD"`
	DBName               string `env:"DB_NAME"`
}

// func checkRequiredEnvVars() Env {
// 	env := Env{}

// 	// pass address of env struct to reflect.ValueOf
// 	v := reflect.ValueOf(&env).Elem()

// 	for i := 0; i < v.NumField(); i++ {
// 		envVar := v.Type().Field(i).Tag.Get("env")
// 		envVarValue := os.Getenv(envVar)

// 		if envVarValue == "" {
// 			panic("Missing required environment variable: " + envVar)
// 		}

// 		// set the value from the environment variable to the struct field
// 		v.Field(i).SetString(envVarValue)
// 	}

// 	return env
// }

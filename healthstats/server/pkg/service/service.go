package service

import (
	"healthstats/pkg/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

const (
	packageName = "service"
)

type ServiceConfig struct {
	AWSRegion            string
	AWSS3RoleArn         string
	AWSAccessKey         string
	AWSSecretAccessKey   string
	AWSS3FilesBucketName string
	Logger               zerolog.Logger
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
}

type Service struct {
	S3Service S3Service
	Logger    zerolog.Logger
	DB        *sqlx.DB
}

func NewService(config ServiceConfig) (*Service, error) {
	s3Service, err := NewS3Service(config)
	if err != nil {
		return nil, err
	}

	pgConfig := model.PostgreSQLDBConfig{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DBName:   config.DBName,
	}

	db, err := model.NewPostgreSQLDB(pgConfig, config.Logger)
	if err != nil {
		return nil, err
	}

	return &Service{
		S3Service: s3Service,
		Logger:    config.Logger,
		DB:        db,
	}, nil
}

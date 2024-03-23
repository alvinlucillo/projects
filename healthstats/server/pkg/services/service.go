package services

import "github.com/rs/zerolog"

const (
	packageName = "services"
)

type ServiceConfig struct {
	AWSRegion            string
	AWSS3RoleArn         string
	AWSAccessKey         string
	AWSSecretAccessKey   string
	AWSS3FilesBucketName string
	Logger               zerolog.Logger
}

type Service struct {
	S3Service S3Service
	Logger    zerolog.Logger
}

func NewService(config ServiceConfig) (*Service, error) {

	s3Service, err := NewS3Service(config)
	if err != nil {
		return nil, err
	}

	return &Service{
		S3Service: s3Service,
		Logger:    config.Logger,
	}, nil
}

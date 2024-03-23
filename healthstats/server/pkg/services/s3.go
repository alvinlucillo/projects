package services

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog"
)

type S3Service interface {
	UploadFile(key string, file multipart.File) (*s3manager.UploadOutput, error)
}

type s3Service struct {
	uploader   *s3manager.Uploader
	bucketName string
	logger     zerolog.Logger
}

func NewS3Service(config ServiceConfig) (S3Service, error) {
	l := config.Logger.With().Str("package", packageName).Str("function", "NewS3Service").Logger()

	// Create a new session in the us-west-2 region.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion)},
	)
	if err != nil {
		fmt.Printf("Error creating session: %s", err.Error())
		l.Err(err).Msg("Error creating session")
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	creds := stscreds.NewCredentials(sess, config.AWSS3RoleArn)

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5MB part size
		u.LeavePartsOnError = true   // Don't delete the parts if the upload fails.
		u.Concurrency = 3            // Download parts concurrently.
		u.S3 = s3.New(sess, &aws.Config{Credentials: creds})
	})

	return s3Service{uploader: uploader, bucketName: config.AWSS3FilesBucketName, logger: config.Logger}, nil
}

func (s s3Service) UploadFile(key string, file multipart.File) (*s3manager.UploadOutput, error) {
	l := s.logger.With().Str("package", packageName).Str("function", "UploadFile").Logger()

	result, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("healthstats-files"),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		// fmt.Printf("Error uploading file: %s", err.Error())
		l.Err(err).Msg("Error uploading file")
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return result, nil
}

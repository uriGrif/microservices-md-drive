package objectstorage

import (
	"bytes"
	"context"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	client *s3.Client
	bucket string
}

func NewS3Storage() ObjectStorage {
	s3Storage := &S3Storage{
		bucket: os.Getenv("S3_BUCKET_NAME"),
	}
	creds := credentials.NewStaticCredentialsProvider(
		os.Getenv("S3_ACCESS_KEY_ID"),
		os.Getenv("S3_SECRET_ACCESS_KEY"),
		"",
	)
	cfg := aws.Config{
		Region:       os.Getenv("S3_REGION"),
		Credentials:  aws.NewCredentialsCache(creds),
		BaseEndpoint: aws.String(os.Getenv("S3_ENDPOINT_URL")),
	}
	s3Storage.client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true // needed for minio and other S3-compatible storages that don't support virtual-hosted style
	})
	if s3Storage.client == nil {
		panic("failed to create S3 client")
	}
	return s3Storage
}

func (s *S3Storage) CreateObject(key string, data []byte) error {
	_, err := s.client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
			Body:   bytes.NewReader(data),
		},
	)
	return err
}

func (s *S3Storage) GetObject(key string) (io.Reader, error) {
	res, err := s.client.GetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (s *S3Storage) GetObjectUrl(key string, expirationSeconds int64) (string, error) {
	presignClient := *s3.NewPresignClient(s.client)
	result, err := presignClient.PresignGetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(expirationSeconds * int64(time.Second))
		},
	)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

func (s *S3Storage) UpdateObject(key string, data []byte) error {
	// same call to PutObject will overwrite the existing object with the new data
	return s.CreateObject(key, data)
}

func (s *S3Storage) DeleteObject(key string) error {
	_, err := s.client.DeleteObject(
		context.Background(),
		&s3.DeleteObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
	)
	return err
}

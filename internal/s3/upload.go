package s3

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/0xdod/uploadtocloud/internal/cloudstorage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Uploader struct {
	client    *s3.Client
	appConfig Config
}

type Config struct {
	Prefix     string `mapstructure:"prefix"`
	BucketName string `mapstructure:"bucketname"`
	Region     string `mapstructure:"region"`
	AccessKey  string `mapstructure:"accesskey"`
	SecretKey  string `mapstructure:"secretkey"`
}

func (s *s3Uploader) Upload(ctx context.Context, data []byte, filename string) (string, error) {
	bucketName := aws.String(s.appConfig.BucketName)
	bucketExists := false

	lsBucketOutput, err := s.client.ListBuckets(ctx, &s3.ListBucketsInput{})

	if err != nil {
		return "", err
	}

	fmt.Printf("S3 buckets: %v\n", lsBucketOutput.Buckets)

	for _, bucket := range lsBucketOutput.Buckets {
		if bucket.Name == bucketName {
			fmt.Printf("Bucket %v exists\n", *bucketName)
			bucketExists = true
			break
		}
	}

	if !bucketExists {
		_, err := s.client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: bucketName,
		})

		if err != nil {
			return "", err
		}
	}

	_, err = s.client.DeletePublicAccessBlock(ctx, &s3.DeletePublicAccessBlockInput{
		Bucket: bucketName,
	})

	if err != nil {
		return "", err
	}

	bucketPolicy := aws.String(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Sid": "PublicReadGetObject",
				"Effect": "Allow",
				"Principal": "*",
				"Action": "s3:GetObject",
				"Resource": "arn:aws:s3:::` + *bucketName + `/*"
			}
		]
	}`)

	_, err = s.client.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
		Bucket: bucketName,
		Policy: bucketPolicy,
	})

	if err != nil {
		return "", err
	}

	if s.appConfig.Prefix != "" {
		filename = s.appConfig.Prefix + "/" + filename
	}

	fmt.Printf("Uploading %v to S3 bucket %v\n", filename, s.appConfig.BucketName)
	uploader := manager.NewUploader(s.client)
	output, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      bucketName,
		Key:         aws.String(filename),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "", err
	}

	fmt.Printf("Successfully uploaded %v to S3\n", filename)
	return output.Location, nil
}

func NewS3Uploader(appCfg Config) cloudstorage.Uploader {
	cred := credentials.NewStaticCredentialsProvider(appCfg.AccessKey, appCfg.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(appCfg.Region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(cred)),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &s3Uploader{client: s3.NewFromConfig(cfg), appConfig: appCfg}
}

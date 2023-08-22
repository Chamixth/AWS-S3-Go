package dao

import (
	dbconfig "aws-s3-go/dbConfig"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateS3BucketDao(bucketName string) (*s3.CreateBucketOutput, error) {
	client, err := dbconfig.CreateS3Client()

	if err != nil {
		return nil, err
	}

	// Create the S3 input parameters
	input := &s3.CreateBucketInput{
		Bucket: &bucketName,
	}

	// Create the bucket
	result, err := client.CreateBucket(context.Background(),input)

	if err != nil{
		return nil,err
	}

	return result,nil
	
}

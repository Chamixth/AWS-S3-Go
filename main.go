package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getS3ObjectURL(bucketName string, objectKey string) (string, error) {
	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Create the input parameters
	input := &s3.GetBucketLocationInput{
		Bucket: &bucketName,
	}

	// Get the bucket location
	output, err := client.GetBucketLocation(context.Background(), input)
	if err != nil {
		return "", err
	}

	// Determine the S3 region based on the bucket location
	region := cfg.Region
	if output.LocationConstraint != types.BucketLocationConstraintUsEast1 {
		region = output.LocationConstraint
	}

	// Construct the S3 object URL
	objectURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, objectKey)

	return objectURL, nil
}

func main() {
	bucketName := "your-bucket-name"
	objectKey := "path/to/your/object.txt"

	objectURL, err := getS3ObjectURL(bucketName, objectKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("S3 Object URL:", objectURL)
}

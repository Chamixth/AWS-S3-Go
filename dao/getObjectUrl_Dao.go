package dao

import (
	"aws-s3-go/dto"
	"fmt"
)

func GetObjectDao(application dto.ObjectUrlRequest) (*string, error) {
	// Construct the S3 object URL using the default region
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", application.BucketName, application.ObjectKey)

	return &objectURL,nil
}

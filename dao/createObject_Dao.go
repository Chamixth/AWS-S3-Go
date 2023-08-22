package dao

import (
	dbconfig "aws-s3-go/dbConfig"
	"aws-s3-go/dto"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateObject(application dto.PutObjectInput) (*s3.PutObjectOutput,error){
	client, err := dbconfig.CreateS3Client()

	if err != nil{
		return nil,err
	}

	input := &s3.PutObjectInput{
		Bucket: &application.Bucket,
		Key: &application.Key,
		Body: application.Body,
	}

	result, err := client.PutObject(context.Background(),input)

	if err != nil{
		return nil,err
	}

	return result,nil
}
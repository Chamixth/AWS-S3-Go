package api

import (
	"aws-s3-go/dao"

	"github.com/gofiber/fiber/v2"
)

func CreateBucketApi(c *fiber.Ctx) error {
	bucketName := c.Query("bucketName")

	result, err := dao.CreateS3BucketDao(bucketName)
	

	if err != nil{
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
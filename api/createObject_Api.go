package api

import (
	"aws-s3-go/dao"
	"aws-s3-go/dto"
	"os"

	"github.com/gofiber/fiber/v2"
)

func CreateObjectApi(c *fiber.Ctx) error {
	var application dto.PutObjectInput

	if err := c.BodyParser(&application); err != nil {
		return err
	}
	
	// Open the local file
	file, err := os.Open(application.LocalFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	application.Body = file // Set the Body field to the opened file

	result, err := dao.CreateObject(application)

	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

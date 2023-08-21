package api

import (
	"aws-s3-go/dao"
	"aws-s3-go/dto"

	"github.com/gofiber/fiber/v2"
)

func GetObjectUrlApi(c *fiber.Ctx) error {
	var application dto.ObjectUrlRequest

	if err := c.BodyParser(&application); err != nil{
		return err
	}

	result, err := dao.GetObjectDao(application)

	if err != nil{
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
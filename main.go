package main

import (
	"aws-s3-go/api"
	"log"

	"github.com/gofiber/fiber/v2"

)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 40 * 1024 * 1024, // Set the body limit to 40MB
	})

	app.Get("/getObjectUrl",api.GetObjectUrlApi)
	app.Post("/createS3Bucket",api.CreateBucketApi)
	app.Post("/createObject",api.CreateObjectApi)
	if err := app.Listen(":4000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
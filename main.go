package main

import (
	"xtest/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(req *models.RequestModels) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"info": "Fiber Validator input Payload",
		})
	})

	app.Post("/auth", func(c *fiber.Ctx) error {
		req := models.RequestModels{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if errs := ValidateStruct(&req); errs != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errs)
		}

		return c.JSON(fiber.Map{
			"message": req,
		})
	})

	app.Listen(":8080")

}

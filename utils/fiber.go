package utils

import "github.com/gofiber/fiber/v2"

func SuccessDataResp(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    data,
	})
}

func ErrorResp(c *fiber.Ctx, err interface{}) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": false,
		"error":   err,
	})
}

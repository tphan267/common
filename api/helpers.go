package api

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResp(c *fiber.Ctx, data interface{}, meta ...ApiResponseMeta) error {
	resp := ApiResponse{
		Success: true,
		Data:    data,
	}
	if len(meta) > 0 {
		resp.Meta = meta[0]
	}
	return c.Status(fiber.StatusOK).JSON(&resp)
}

func ErrorResp(c *fiber.Ctx, err ApiError, meta ...ApiResponseMeta) error {
	resp := ApiResponse{
		Success: true,
		Error:   err,
	}
	if len(meta) > 0 {
		resp.Meta = meta[0]
	}
	return c.Status(fiber.StatusOK).JSON(&resp)
}

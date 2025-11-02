package utils

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Code    int         `json:"code" example:"200"`
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Request successfully processed"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int         `json:"code" example:"400"`
	Status  string      `json:"status" example:"error"`
	Message string      `json:"message" example:"Invalid request parameters"`
	Error   interface{} `json:"error"`
}

func NewSuccessResponse(send *fiber.Ctx, code int, message string, data interface{}) error {
	response := SuccessResponse{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    data,
	}

	return send.Status(code).JSON(response)
}

func NewErrorResponse(send *fiber.Ctx, code int, message string, err interface{}) error {
	response := ErrorResponse{
		Code:    code,
		Status:  "error",
		Message: message,
		Error:   err,
	}

	return send.Status(code).JSON(response)
}

func OKResponse(ctx *fiber.Ctx, message string, data interface{}) error {
	return NewSuccessResponse(ctx, 200, message, data)
}

func CreatedResponse(ctx *fiber.Ctx, message string, data interface{}) error {
	return NewSuccessResponse(ctx, 201, message, data)
}

func NoContentResponse(ctx *fiber.Ctx, message string) error {
	return NewSuccessResponse(ctx, 200, message, nil)
}

func BadRequestResponse(ctx *fiber.Ctx, message string, err interface{}) error {
	return NewErrorResponse(ctx, 400, message, err)
}

func NotFoundResponse(ctx *fiber.Ctx, message string, err interface{}) error {
	return NewErrorResponse(ctx, 404, message, err)
}

func InternalServerErrorResponse(ctx *fiber.Ctx, message string, err interface{}) error {
	return NewErrorResponse(ctx, 500, message, err)
}

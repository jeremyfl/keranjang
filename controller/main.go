package controller

import "github.com/labstack/echo/v4"

// GenerateResponse generating the HTTP API response
func GenerateResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, map[string]interface{}{
		"status_code": statusCode,
		"message":     message,
		"data":        data,
	})
}

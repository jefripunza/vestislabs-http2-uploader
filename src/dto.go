package src

import "github.com/labstack/echo/v4"

type JSONResponse struct {
	StatusCode int    `json:"StatusCode"`
	Success    bool   `json:"Success"`
	Message    string `json:"Message"`
}

func SendJSONResponse(c echo.Context, statusCode int, success bool, message string) error {
	resp := JSONResponse{
		StatusCode: statusCode,
		Success:    success,
		Message:    message,
	}
	return c.JSON(statusCode, resp)
}

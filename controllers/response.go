package controllers

import "github.com/labstack/echo/v4"

type Response struct {
	Data interface{} `json:"data"`
}

func Success(c echo.Context, data interface{}) error {
	res := Response{
		Data: data,
	}

	return c.JSON(200, res)
}

func Error(c echo.Context, status int) error {
	res := Response{
		Data: nil,
	}

	return c.JSON(status, res)
}

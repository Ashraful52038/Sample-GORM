package handlers

import (
	"EchoV4/models/request"
	"EchoV4/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {

	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err.Error()))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Validation Failed", err.Error()))
	}

	user := request.User{
		Name:  req.Name,
		Email: req.Email,
	}

	return c.JSON(201, response.SuccessResponse("User created successfully", user))
}

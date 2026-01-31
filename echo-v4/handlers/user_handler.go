package handlers

import (
	"EchoV4/models/request"
	"EchoV4/models/response"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {

	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err.Error()))
	}

	user := request.User{
		Id:    strconv.FormatInt(time.Now().Unix(), 10),
		Name:  req.Name,
		Email: req.Email,
		Age:   25,
	}

	return c.JSON(201, response.SuccessResponse("User created successfully", user))
}

// GetUser Single user - GET /users/:id
func GetUser(c echo.Context) error {
	id := c.Param("id")

	user := request.User{
		Id:    id,
		Name:  "Ashraful",
		Email: "xyz@mail.com",
		Age:   30,
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("find user", user))
}

// UpdateUser - PUT /users/:id
func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var req request.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err.Error()))
	}

	// Update logic
	updatedUser := request.User{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
		Age:   25, // Default age
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("User updated successfully", updatedUser))
}

// Delete user
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, response.SuccessResponse("User Deleted Successfully", map[string]string{
		"id": id}))
}

// Get All user
func GetAllUsers(c echo.Context) error {
	users := []request.User{
		{
			Id:    "1",
			Name:  "Ashraful",
			Email: "ash@mail.com",
			Age:   25,
		},
		{
			Id:    "2",
			Name:  "Rahim",
			Email: "rahim@mail.com",
			Age:   29,
		},
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("Users Retrived Successfully", users))
}

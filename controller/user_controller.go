package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-app/controller/request"
	"product-app/controller/response"
	"product-app/middlewares"
	"product-app/service"
	"strconv"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/api/v1/users")
	g.Use(middlewares.RequireAuth())
	g.GET("", userController.GetAllUsers)
	g.GET("/:id", userController.GetUserById)
	g.POST("", userController.AddUser)
	g.PUT("", userController.UpdateUser)
	g.DELETE("/:id", userController.DeleteUser)

}

func (userController *UserController) GetUserById(c echo.Context) error {
	id := c.Param("id")

	userId, _ := strconv.Atoi(id)

	user, err := userController.userService.GetUserById(uint(userId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	responseUser := response.GetUserByIdResponse{}.ToResponse(&user)

	return c.JSON(http.StatusOK, responseUser)
}

func (userController *UserController) GetAllUsers(c echo.Context) error {
	users := userController.userService.GetAllUsers()
	responseUsers := make([]response.GetAllUsersResponse, 0)

	for _, user := range users {
		responseUsers = append(responseUsers, response.GetAllUsersResponse{}.ToResponse(&user))
	}

	return c.JSON(http.StatusOK, responseUsers)
}

func (userController *UserController) AddUser(c echo.Context) error {
	var addUserRequest request.AddUserRequest

	err := c.Bind(&addUserRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	err = userController.userService.AddUser(addUserRequest.ToModel())

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{Message: "User added successfully"})
}

func (userController *UserController) UpdateUser(c echo.Context) error {
	var updateUserRequest request.UpdateUserRequest

	err := c.Bind(&updateUserRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	err = userController.userService.UpdateUser(updateUserRequest.ToModel())

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Message: "User updated successfully"})
}

func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	userId, _ := strconv.Atoi(id)

	err := userController.userService.DeleteUser(uint(userId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Message: "User deleted successfully"})
}

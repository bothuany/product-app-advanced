package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-app/controller/request"
	"product-app/controller/response"
	"product-app/middlewares"
	"product-app/service"
	"time"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (authController *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/v1/auth/login", authController.Login)
	e.POST("/api/v1/auth/register", authController.Register)

	g := e.Group("/api/v1/auth/validate")
	g.Use(middlewares.RequireAuth())
	g.GET("", authController.Validate)
}

func (authController *AuthController) Login(c echo.Context) error {
	var loginRequest request.LoginRequest

	err := c.Bind(&loginRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	token, authErr := authController.authService.Login(loginRequest.ToModel())

	if authErr != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{ErrorDescription: authErr.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour * 30),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusOK, response.LoginResponse{Token: token})
}

func (authController *AuthController) Register(c echo.Context) error {
	var registerRequest request.RegisterRequest

	err := c.Bind(&registerRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	token, authErr := authController.authService.Register(registerRequest.ToModel())

	if authErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: authErr.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour * 30),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusCreated, response.LoginResponse{Token: token})
}

func (authController *AuthController) Validate(c echo.Context) error {
	user := c.Get("user")
	return c.JSON(http.StatusOK, user)
}

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"product-app/controller"
	"product-app/initializers"
	"product-app/persistence"
	"product-app/service"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	e := echo.New()

	//configurationManager := app.NewConfigurationManager()

	productRepository := persistence.NewProductRepository()
	userRepository := persistence.NewUserRepository()

	productService := service.NewProductService(productRepository)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)

	productController := controller.NewProductController(productService)
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)

	productController.RegisterRoutes(e)
	authController.RegisterRoutes(e)
	userController.RegisterRoutes(e)

	port := os.Getenv("PORT")
	fmt.Println("Server is running on port: ", port)
	e.Start("localhost:" + port)
}

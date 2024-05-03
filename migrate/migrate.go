package main

import (
	"product-app/domain"
	"product-app/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&domain.Product{})
	initializers.DB.AutoMigrate(&domain.User{})
}

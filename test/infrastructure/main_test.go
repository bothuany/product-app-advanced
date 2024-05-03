package infrastructure

import (
	"fmt"
	"os"
	"product-app/initializers"
	"product-app/persistence"
	"testing"
)

var productRepository persistence.IProductRepository
var userRepository persistence.IUserRepository

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	userRepository = persistence.NewUserRepository()
	productRepository = persistence.NewProductRepository()
}

func TestMain(m *testing.M) {
	fmt.Println("-------------Before all tests-------------")
	exitCode := m.Run()
	fmt.Println("-------------After all tests-------------")
	os.Exit(exitCode)
}

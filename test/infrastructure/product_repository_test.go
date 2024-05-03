package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"product-app/domain"
	"product-app/initializers"
	"testing"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func setupProducts() {
	TestDataInitialize()
}

func clearProducts() {
	TruncateTestData()
}

func TestGetAllProducts(t *testing.T) {
	setupProducts()
	expectedProducts := []domain.Product{
		{Model: gorm.Model{ID: 1}, Name: "AirFryer", Price: 3000.0, Discount: 22.0, Store: "ABC TECH"},
		{Model: gorm.Model{ID: 2}, Name: "Ütü", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Model: gorm.Model{ID: 3}, Name: "Çamaşır Makinesi", Price: 10000.0, Discount: 15.0, Store: "ABC TECH"},
		{Model: gorm.Model{ID: 4}, Name: "Lambader", Price: 2000.0, Discount: 0.0, Store: "Dekorasyon Sarayı"},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clearProducts()

}

func TestGetAllProductsByStore(t *testing.T) {
	setupProducts()
	expectedProducts := []domain.Product{
		{Model: gorm.Model{ID: 1}, Name: "AirFryer", Price: 3000.0, Discount: 22.0, Store: "ABC TECH"},
		{Model: gorm.Model{ID: 2}, Name: "Ütü", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Model: gorm.Model{ID: 3}, Name: "Çamaşır Makinesi", Price: 10000.0, Discount: 15.0, Store: "ABC TECH"},
	}

	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clearProducts()
}

func TestAddProduct(t *testing.T) {
	expectedProducts := []domain.Product{
		{Model: gorm.Model{ID: 1}, Name: "Kupa", Price: 100.0, Discount: 0.0, Store: "RBD"},
	}

	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    100.0,
		Discount: 0.0,
		Store:    "RBD",
	}
	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		actualProducts := productRepository.GetAllProducts()
		actualProducts[0].Model = gorm.Model{ID: uint(len(actualProducts))}
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clearProducts()
}

func TestGetProductById(t *testing.T) {
	setupProducts()
	expectedProduct := domain.Product{Model: gorm.Model{ID: 1}, Name: "AirFryer", Price: 3000.0, Discount: 22.0, Store: "ABC TECH"}
	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetProductById(1)
		_, err := productRepository.GetProductById(5)
		assert.Equal(t, expectedProduct, actualProduct)
		assert.Equal(t, "Product with id 5 not found", err.Error())
	})
	clearProducts()
}

func TestDeleteProductById(t *testing.T) {
	setupProducts()

	t.Run("DeleteProductById", func(t *testing.T) {
		productRepository.DeleteProductById(4)
		_, err := productRepository.GetProductById(4)
		assert.Equal(t, "Product with id 4 not found", err.Error())
	})

	clearProducts()
}

func TestUpdateProduct(t *testing.T) {
	setupProducts()
	expectedProduct := domain.Product{Model: gorm.Model{ID: 1}, Name: "Fırın", Price: 4000.0, Discount: 22.0, Store: "ABC TECH"}
	updatedProduct := domain.Product{Model: gorm.Model{ID: 1}, Name: "Fırın", Price: 4000.0, Discount: 22.0, Store: "ABC TECH"}

	t.Run("UpdateProduct", func(t *testing.T) {
		productBeforeUpdate, _ := productRepository.GetProductById(1)
		assert.Equal(t, "AirFryer", productBeforeUpdate.Name)
		assert.Equal(t, float32(3000.0), productBeforeUpdate.Price)
		productRepository.UpdateProduct(updatedProduct)
		productAfterUpdate, _ := productRepository.GetProductById(1)
		assert.Equal(t, expectedProduct, productAfterUpdate)
	})

	clearProducts()
}

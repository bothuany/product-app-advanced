package persistence

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"product-app/domain"
	"product-app/initializers"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetProductById(productId uint) (domain.Product, error)
	DeleteProductById(productId uint) error
	UpdateProduct(product domain.Product) error
}

type ProductRepository struct {
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	var products []domain.Product
	result := initializers.DB.Find(&products)

	if result.Error != nil {
		log.Error("Error while fetching products: %v\n", result.Error)
		return []domain.Product{}
	}

	return products
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	var products []domain.Product
	result := initializers.DB.Where("store = ?", storeName).Find(&products)

	if result.Error != nil {
		log.Error("Error while fetching products: %v\n", result.Error)
		return []domain.Product{}
	}

	return products
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		log.Error("Error while adding product: %v\n", result.Error)
		return result.Error
	}

	log.Info("Product added with id: %v\n", product.ID)

	return nil
}

func (productRepository *ProductRepository) GetProductById(productId uint) (domain.Product, error) {
	var product domain.Product
	result := initializers.DB.First(&product, productId)

	if result.Error != nil {
		log.Error("Error while fetching product: %v\n", result.Error)
		return domain.Product{}, errors.New(fmt.Sprintf("Product with id %d not found", productId))
	}

	return product, nil
}

func (productRepository *ProductRepository) DeleteProductById(productId uint) error {
	initializers.DB.Delete(&domain.Product{}, productId)

	return nil
}

func (productRepository *ProductRepository) UpdateProduct(product domain.Product) error {
	err := productRepository.checkProductExists(product.ID)

	if err != nil {
		return err
	}

	// Start a new DB transaction
	tx := initializers.DB.Begin()

	// Rollback transaction in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Commit transaction if no error
	if err := tx.Error; err != nil {
		return err
	}

	// Update each field if it's not the zero value
	if product.Name != "" {
		tx.Model(&domain.Product{}).Where("id = ?", product.ID).UpdateColumn("name", product.Name)
	}
	if product.Price != 0 {
		tx.Model(&domain.Product{}).Where("id = ?", product.ID).UpdateColumn("price", product.Price)
	}
	if product.Discount != 0 {
		tx.Model(&domain.Product{}).Where("id = ?", product.ID).UpdateColumn("discount", product.Discount)
	}
	if product.Store != "" {
		tx.Model(&domain.Product{}).Where("id = ?", product.ID).UpdateColumn("store", product.Store)
	}

	// Commit the transaction
	tx.Commit()

	log.Info("Product updated with id: %v\n", product.ID)

	return nil
}
func (productRepository *ProductRepository) checkProductExists(productId uint) error {
	_, getErr := productRepository.GetProductById(productId)

	if getErr != nil {
		log.Error("Error while checking product: %v\n", getErr)
		return errors.New(fmt.Sprintf("Product with id %d not found", productId))
	}

	return nil
}

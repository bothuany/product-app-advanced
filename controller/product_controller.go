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

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products", productController.GetAllProducts)

	g := e.Group("/api/v1/products")
	g.Use(middlewares.RequireAuth())
	g.POST("", productController.AddProduct)
	g.PUT("", productController.UpdateProduct)
	g.DELETE("/:id", productController.DeleteProductById)
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	id := c.Param("id")
	productId, _ := strconv.Atoi(id)

	product, err := productController.productService.GetProductById(uint(productId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	responseProduct := response.GetProductByIdResponse{}.ToResponse(&product)

	return c.JSON(http.StatusOK, responseProduct)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParams().Get("store")
	if store != "" {
		productsByStore := productController.productService.GetAllProductsByStore(store)
		responseProducts := make([]response.GetAllProductsResponse, 0)

		for _, product := range productsByStore {
			responseProducts = append(responseProducts, response.GetAllProductsResponse{}.ToResponse(&product))
		}

		return c.JSON(http.StatusOK, responseProducts)
	}

	products := productController.productService.GetAllProducts()
	responseProducts := make([]response.GetAllProductsResponse, 0)

	for _, product := range products {
		responseProducts = append(responseProducts, response.GetAllProductsResponse{}.ToResponse(&product))
	}

	return c.JSON(http.StatusOK, responseProducts)
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest

	err := c.Bind(&addProductRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	err = productController.productService.AddProduct(addProductRequest.ToModel())

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{Message: "Product added successfully"})

}

func (productController *ProductController) UpdateProduct(c echo.Context) error {
	var updateProductRequest request.UpdateProductRequest

	err := c.Bind(&updateProductRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	err = productController.productService.UpdateProduct(updateProductRequest.ToModel())

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Message: "Product updated successfully"})
}

func (productController *ProductController) DeleteProductById(c echo.Context) error {
	id := c.Param("id")
	productId, _ := strconv.Atoi(id)

	err := productController.productService.DeleteProductById(uint(productId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Message: "Product deleted successfully"})
}

package response

import (
	"product-app/domain"
)

type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type GetProductByIdResponse struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (productResponse GetProductByIdResponse) ToResponse(product *domain.Product) GetProductByIdResponse {
	return GetProductByIdResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

type GetAllProductsResponse struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (productResponse GetAllProductsResponse) ToResponse(product *domain.Product) GetAllProductsResponse {
	return GetAllProductsResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserByIdResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func (userResponse GetUserByIdResponse) ToResponse(user *domain.User) GetUserByIdResponse {
	return GetUserByIdResponse{
		Id:    user.ID,
		Email: user.Email,
	}
}

type GetAllUsersResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func (userResponse GetAllUsersResponse) ToResponse(user *domain.User) GetAllUsersResponse {
	return GetAllUsersResponse{
		Id:    user.ID,
		Email: user.Email,
	}
}

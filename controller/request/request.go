package request

import "product-app/service/model"

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (addProductRequest AddProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}

type UpdateProductRequest struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (updateProductRequest UpdateProductRequest) ToModel() model.ProductUpdate {
	return model.ProductUpdate{
		Id:       updateProductRequest.Id,
		Name:     updateProductRequest.Name,
		Price:    updateProductRequest.Price,
		Discount: updateProductRequest.Discount,
		Store:    updateProductRequest.Store,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (registerRequest RegisterRequest) ToModel() model.UserCreate {
	return model.UserCreate{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginRequest LoginRequest) ToModel() model.UserLogin {
	return model.UserLogin{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}

type UpdateUserRequest struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (updateUserRequest UpdateUserRequest) ToModel() model.UserUpdate {
	return model.UserUpdate{
		Id:       updateUserRequest.Id,
		Email:    updateUserRequest.Email,
		Password: updateUserRequest.Password,
	}
}

type AddUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (addUserRequest AddUserRequest) ToModel() model.UserCreate {
	return model.UserCreate{
		Email:    addUserRequest.Email,
		Password: addUserRequest.Password,
	}
}

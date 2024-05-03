package model

type ProductCreate struct {
	Name     string
	Price    float32
	Discount float32
	Store    string
}

type ProductUpdate struct {
	Id       uint
	Name     string
	Price    float32
	Discount float32
	Store    string
}

type UserCreate struct {
	Email    string
	Password string
}

type UserLogin struct {
	Email    string
	Password string
}

type UserUpdate struct {
	Id       uint
	Email    string
	Password string
}

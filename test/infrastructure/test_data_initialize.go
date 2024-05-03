package infrastructure

import "product-app/initializers"

var INSERT_PRODUCTS = `
INSERT INTO products (name, price, discount,store) 
VALUES('AirFryer',3000.0, 22.0, 'ABC TECH'),
('Ütü',1500.0, 10.0, 'ABC TECH'),
('Çamaşır Makinesi',10000.0, 15.0, 'ABC TECH'),
('Lambader',2000.0, 0.0, 'Dekorasyon Sarayı');
`

var INSERT_USERS = `
INSERT INTO users (email, password)
VALUES('test@mail.com', 'test'),
('test2@mail.com', 'test2');
`

func TestDataInitialize() {
	initializers.DB.Exec(INSERT_PRODUCTS)
	initializers.DB.Exec(INSERT_USERS)
}

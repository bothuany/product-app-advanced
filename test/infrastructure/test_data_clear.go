package infrastructure

import "product-app/initializers"

func TruncateTestData() {
	initializers.DB.Exec("TRUNCATE TABLE products RESTART IDENTITY")
	initializers.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY")
}

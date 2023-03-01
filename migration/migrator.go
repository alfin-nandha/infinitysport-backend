package migration

import (
	Cart "project/e-comerce/features/carts/data"
	Order "project/e-comerce/features/orders/data"
	Product "project/e-comerce/features/products/data"
	User "project/e-comerce/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(User.User{})
	db.AutoMigrate(Product.Product{})
	db.AutoMigrate(Cart.Cart{})
	db.AutoMigrate(Order.Order{})
	db.AutoMigrate(Order.OrderDetail{})
	db.AutoMigrate(Order.Address{})
	db.AutoMigrate(Order.Payment{})
}

package factory

import (
	UserBusiness "project/e-comerce/features/users/business"
	UserData "project/e-comerce/features/users/data"
	UserPresentation "project/e-comerce/features/users/presentation"

	AuthBusiness "project/e-comerce/features/auth/business"
	AuthData "project/e-comerce/features/auth/data"
	AuthPresentation "project/e-comerce/features/auth/presentation"

	ProductBusiness "project/e-comerce/features/products/bussiness"
	ProductData "project/e-comerce/features/products/data"
	ProductPresentation "project/e-comerce/features/products/presentation"

	CartBusiness "project/e-comerce/features/carts/business"
	CartData "project/e-comerce/features/carts/data"
	CartPresentation "project/e-comerce/features/carts/presentation"

	OrderBusiness "project/e-comerce/features/orders/bussiness"
	OrderData "project/e-comerce/features/orders/data"
	OrderPresentation "project/e-comerce/features/orders/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *UserPresentation.UserHandler
	AuthPresenter    *AuthPresentation.AuthHandler
	ProductPresenter *ProductPresentation.ProductHandler
	OrderPresenter   *OrderPresentation.OrderHandler
	CartPresenter    *CartPresentation.CartHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := UserData.NewUserRepository(dbConn)
	userBusiness := UserBusiness.NewUserBusiness(userData)
	userPresentation := UserPresentation.NewUserHandler(userBusiness)

	authData := AuthData.NewAuthRepository(dbConn)
	authBusiness := AuthBusiness.NewAuthBusiness(authData)
	authPresentation := AuthPresentation.NewAuthHandler(authBusiness)

	productData := ProductData.NewProductRepository(dbConn)
	productBusiness := ProductBusiness.NewProductBusiness(productData)
	productPresentation := ProductPresentation.NewProductHandler(productBusiness)

	cartData := CartData.NewCartRepository(dbConn)
	cartBusiness := CartBusiness.NewCartBusiness(cartData)
	cartPresentation := CartPresentation.NewCartHandler(cartBusiness)

	orderData := OrderData.NewOrderRepository(dbConn)
	orderBusiness := OrderBusiness.NewOrderBusiness(orderData, cartData)
	orderPresentation := OrderPresentation.NewOrderHandler(orderBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		AuthPresenter:    authPresentation,
		ProductPresenter: productPresentation,
		CartPresenter:    cartPresentation,
		OrderPresenter:   orderPresentation,
	}
}

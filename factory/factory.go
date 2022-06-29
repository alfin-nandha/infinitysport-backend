package factory

import (
	_userBusiness "project/e-comerce/features/users/business"
	_userData "project/e-comerce/features/users/data"
	_userPresentation "project/e-comerce/features/users/presentation"

	_authBusiness "project/e-comerce/features/auth/business"
	_authData "project/e-comerce/features/auth/data"
	_authPresentation "project/e-comerce/features/auth/presentation"

	_productBusiness "project/e-comerce/features/products/bussiness"
	_productData "project/e-comerce/features/products/data"
	_productPresentation "project/e-comerce/features/products/presentation"

	_orderBusiness "project/e-comerce/features/orders/bussiness"
	_orderData "project/e-comerce/features/orders/data"
	_orderPresentation "project/e-comerce/features/orders/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
	AuthPresenter *_authPresentation.AuthHandler
	ProductPresenter *_productPresentation.ProductHandler
	OrderPresenter *_orderPresentation.OrderHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	productData := _productData.NewProductRepository(dbConn)
	productBusiness := _productBusiness.NewProductBusiness(productData)
	productPresentation := _productPresentation.NewProductHandler(productBusiness)

	orderData := _orderData.NewOrderRepository(dbConn)
	// tambah parameter cartData
	orderBusiness := _orderBusiness.NewOrderBusiness(orderData)
	orderPresentation := _orderPresentation.NewOrderHandler(orderBusiness)


	return Presenter{
		UserPresenter: userPresentation,
		AuthPresenter: authPresentation,
		ProductPresenter: productPresentation,
		OrderPresenter: orderPresentation,		
	}
}

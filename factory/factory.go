package factory

import (
	_userBusiness "project/e-comerce/features/users/business"
	_userData "project/e-comerce/features/users/data"
	_userPresentation "project/e-comerce/features/users/presentation"

	_authBusiness "project/e-comerce/features/auth/business"
	_authData "project/e-comerce/features/auth/data"
	_authPresentation "project/e-comerce/features/auth/presentation"

	"gorm.io/gorm"
)

type Presenter struct{
	UserPresenter *_userPresentation.UserHandler
	AuthPresenter *_authPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter{

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewUserRepository(dbConn)
	authBusiness := _authBusiness.NewUserBusiness(authData)
	authPresentation := _authPresentation.NewUserHandler(authBusiness)

	return Presenter{
		UserPresenter: userPresentation,
		AuthPresenter: authPresentation,
		
	}
}
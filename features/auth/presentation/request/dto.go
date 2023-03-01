package request

import (
	Auth "project/e-comerce/features/auth"
)

type User struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func ToCore(userReq User) Auth.Core {
	userCore := Auth.Core{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}

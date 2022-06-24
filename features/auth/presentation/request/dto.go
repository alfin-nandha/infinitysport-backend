package request

import (
	"project/e-comerce/features/auth"
)

type User struct{
	Email     string    `json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

func ToCore(userReq User) (auth.Core){
	userCore := auth.Core{
		Email: userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
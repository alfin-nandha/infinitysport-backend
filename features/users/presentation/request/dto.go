package request

import "project/e-comerce/features/users"

type User struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq User) users.Core {
	userCore := users.Core{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	
	}
	return userCore
}

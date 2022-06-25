package request

import "project/e-comerce/features/users"

type User struct{
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
	Role_ID    int 		`json:"role_id" form:"role_id"`
}

func ToCore(userReq User) (users.Core){
	userCore := users.Core{
		Name: userReq.Name,
		Email: userReq.Email,
		Password: userReq.Password,
	
	}
	return userCore
}
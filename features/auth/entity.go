package auth

import (
	"project/e-comerce/features/users/data"
)


type Core struct{
	Email 		string 
	Password 	string
}

type Business interface{
	Login(Core)(token string,Name string, err error)
}

type Data interface{
	FindUser(email string)(data.User , error)
}
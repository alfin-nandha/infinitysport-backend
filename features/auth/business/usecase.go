package business

import (
	"project/e-comerce/features/auth"
	"project/e-comerce/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct{
	userData auth.Data
}

func NewUserBusiness(usrData auth.Data) auth.Business {
	return &authUseCase{
		userData: usrData,
	}
}

func (uc *authUseCase) Login(param auth.Core)(string, string, error){
	response, errFind := uc.userData.FindUser(param.Email)
	if errFind != nil{
		return "", "",errFind
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(param.Password))
	if errCompare != nil{
		return "", "",errCompare
	}
	token,err := middlewares.CreateToken(int(response.ID))
	
	return token,response.Name,err
}

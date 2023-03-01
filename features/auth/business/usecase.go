package business

import (
	Auth "project/e-comerce/features/auth"
	Middlewares "project/e-comerce/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userData Auth.Data
}

func NewAuthBusiness(usrData Auth.Data) Auth.Business {
	return &authUseCase{
		userData: usrData,
	}
}

func (uc *authUseCase) Login(data Auth.Core) (string, string, error) {
	response, errFind := uc.userData.FindUser(data.Email)
	if errFind != nil {
		return "", "", errFind
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(data.Password))
	if errCompare != nil {
		return "", "", errCompare
	}
	token, err := Middlewares.CreateToken(int(response.ID))

	return token, response.Name, err
}

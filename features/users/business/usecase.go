package business

import (
	"errors"
	Users "project/e-comerce/features/users"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData Users.Data
}

func NewUserBusiness(usrData Users.Data) Users.Business {
	return &userUseCase{
		userData: usrData,
	}
}

func (uc *userUseCase) GetAllData(limit, offset int) (response []Users.Core, err error) {
	resp, errData := uc.userData.SelectData(limit, offset)
	return resp, errData
}

func (uc *userUseCase) GetDataById(id int) (response Users.Core, err error) {
	resp, errData := uc.userData.SelectDataById(id)
	return resp, errData
}

func (uc *userUseCase) InsertData(userRequest Users.Core) (row int, err error) {

	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" {
		return -1, errors.New("all data must be filled")
	}

	passWillBcrypt := []byte(userRequest.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
		return -2, errors.New("hashing password failed")
	}
	userRequest.Password = string(hash)
	result, err := uc.userData.InsertData(userRequest)
	if err != nil {
		return 0, errors.New("failed to insert data")
	}
	return result, nil
}

func (uc *userUseCase) DeleteData(id int) (row int, err error) {
	result, err := uc.userData.DeleteData(id)
	if err != nil {
		return 0, errors.New("no data")
	}
	return result, err
}

func (uc *userUseCase) UpdateData(userReq Users.Core, id int) (row int, err error) {
	updateMap := make(map[string]interface{})
	if userReq.Name != "" {
		updateMap["name"] = &userReq.Name
	}
	if userReq.Email != "" {
		updateMap["email"] = &userReq.Email
	}
	if userReq.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return -1, errors.New("hasing password failed")
		}
		updateMap["password"] = &hash
	}

	result, err := uc.userData.UpdateData(updateMap, id)
	if err != nil {
		return 0, errors.New("no data user for updated")
	}
	return result, nil
}

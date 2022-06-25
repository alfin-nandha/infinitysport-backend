package business

import (
	"errors"
	"project/e-comerce/features/users"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct{
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUseCase{
		userData: usrData,
	}
}

func (uc *userUseCase) GetAllData(param string)(response []users.Core,err error){
	resp,errData := uc.userData.SelectData(param)
	return resp,errData
}

func (uc *userUseCase) GetDataById(id int)(response users.Core, err error){
	response, err = uc.userData.SelectDataById(id)
	return response,err
}

func (uc *userUseCase) InsertData(userRequest users.Core)(error){
	
	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == ""{
		return errors.New("all data must be filled")
	}


	passWillBcrypt := []byte(userRequest.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
        return errors.New("hashing password failed")
    }
	userRequest.Password = string(hash)
	err := uc.userData.InsertData(userRequest)
	return err
}

func (uc *userUseCase) DeleteData(id int)(err error){
	err = uc.userData.DeleteData(id)
	return err
}

func (uc *userUseCase) UpdateData(userReq users.Core, id int)(err error){
	updateMap := make(map[string]interface{})
	if userReq.Name != ""{
		updateMap["name"] = &userReq.Name
	}
	if userReq.Email != ""{
		updateMap["email"] = &userReq.Email
	}
	if userReq.Password != ""{
		hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil{
			return errors.New("hasing password failed")
		}
		updateMap["password"] = &hash
	}

	err = uc.userData.UpdateData(updateMap, id)
	return err
}

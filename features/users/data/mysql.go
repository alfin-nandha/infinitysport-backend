package data

import (
	"errors"
	"project/e-comerce/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct{
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository)SelectData(data string) (response []users.Core, err error){
	var dataUser []User
	result := repo.db.Preload("Role").Find(&dataUser)
	if result.Error != nil{
		return []users.Core{}, result.Error
	}
	// fmt.Println(dataUser[6].Role.Role_name)
	return toCoreList(dataUser),result.Error
}


func (repo *mysqlUserRepository)SelectDataById(id int) (response users.Core, err error){
	datauser := User{}
	result := repo.db.Find(&datauser, id)
	if result.Error != nil{
		return users.Core{}, result.Error
	}
	return toCore(datauser), err
}

func (repo *mysqlUserRepository)InsertData(userData users.Core)error{
	userModel := fromCore(userData)
	result := repo.db.Create(&userModel)
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("failed insert data")
	}
	return  nil
}

func (repo *mysqlUserRepository)DeleteData(id int) (error){
	datauser := User{}
	err := repo.db.Delete(&datauser,id)
	return err.Error
}

func (repo *mysqlUserRepository)UpdateData(dataReq map[string]interface{}, id int)(error){

	model := User{}
	model.ID = uint(id)
	result := repo.db.Model(model).Updates(dataReq)
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("failed update data")
	}
	return  nil

}
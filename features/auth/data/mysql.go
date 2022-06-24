package data

import (
	"project/e-comerce/features/auth"
	"project/e-comerce/features/users/data"

	"gorm.io/gorm"
)

type mysqlUserRepository struct{
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) auth.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository)FindUser(email string) (response data.User, err error){
	datauser := data.User{}
	result := repo.db.Where("email = ?", email).Find(&datauser)
	if result.Error != nil{
		return data.User{}, result.Error
	}
	return datauser, err
}
package data

import (
	Auth "project/e-comerce/features/auth"

	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) Auth.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}

func (repo *mysqlAuthRepository) FindUser(email string) (response Auth.Core, err error) {
	dataUser := User{}
	result := repo.db.Where("email = ?", email).Find(&dataUser)
	if result.Error != nil {
		return Auth.Core{}, result.Error
	}

	return dataUser.toCore(), nil
}

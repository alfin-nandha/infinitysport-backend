package data

import (
	"project/e-comerce/features/auth"

	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}

func (repo *mysqlAuthRepository) FindUser(data auth.Core) (response []string, err error) {
	dataUser := authCore(data)
	resultEmail := repo.db.Where("email = ?", dataUser.Email).Find(&dataUser)
	if resultEmail.Error != nil {
		return []string{}, resultEmail.Error
	} else {
		resultPassword := repo.db.Where("password = ?", dataUser.Password).Find(&dataUser)
		if resultPassword != nil {
			return []string{}, resultPassword.Error
		}
		return resultPassword, nil
	}
	return resultEmail, nil
}

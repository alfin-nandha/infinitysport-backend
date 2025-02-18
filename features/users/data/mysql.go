package data

import (
	"errors"
	Users "project/e-comerce/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) Users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) SelectData(limit, offset int) (response []Users.Core, err error) {
	var dataUser []User
	result := repo.db.Find(&dataUser)
	if result.Error != nil {
		return []Users.Core{}, result.Error
	}
	return toCoreList(dataUser), result.Error
}

func (repo *mysqlUserRepository) SelectDataById(id int) (response Users.Core, err error) {
	datauser := User{}
	result := repo.db.Find(&datauser, id)
	if result.Error != nil {
		return Users.Core{}, result.Error
	}
	return toCore(datauser), nil
}

func (repo *mysqlUserRepository) InsertData(userData Users.Core) (row int, err error) {
	userModel := fromCore(userData)
	result := repo.db.Create(&userModel)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) DeleteData(id int) (row int, err error) {
	datauser := User{}
	result := repo.db.Delete(&datauser, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) UpdateData(dataReq map[string]interface{}, id int) (row int, err error) {
	model := User{}
	model.ID = uint(id)
	result := repo.db.Model(model).Updates(dataReq)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}
	return int(result.RowsAffected), nil
}

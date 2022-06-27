package data

import (
	"errors"
	"project/e-comerce/features/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct{
	db *gorm.DB
}

func NewProductRepository(conn *gorm.DB) products.Data {
	return &mysqlProductRepository{
		db: conn,
	}
}

func (repo *mysqlProductRepository)SelectData() (response []products.Core, err error){
	var dataProduct []Product
	result := repo.db.Find(&dataProduct)
	if result.Error != nil{
		return []products.Core{}, result.Error
	}
	return ToCoreList(dataProduct),result.Error
}


func (repo *mysqlProductRepository)SelectDataByID(id int) (response products.Core, err error){
	dataProduct := Product{}
	result := repo.db.Find(&dataProduct, id)
	if result.Error != nil{
		
		return products.Core{}, result.Error
	}
	
	return toCore(dataProduct), err
}

func (repo *mysqlProductRepository)InsertData(ProductData products.Core)error{
	ProductModel := fromCore(ProductData)
	result := repo.db.Create(&ProductModel)
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("failed insert data")
	}
	return  nil
}

func (repo *mysqlProductRepository)DeleteDataByID(id int, userId int) (error){
	dataProduct := Product{}
	result := repo.db.Where("user_id = ?", userId).Delete(&dataProduct,id)
	if result.RowsAffected == 0{
		return errors.New("no rows affected")
	}
	return result.Error
}

func (repo *mysqlProductRepository)UpdateDataByID(dataReq map[string]interface{}, id int, userId int)(error){

	model := Product{}
	model.ID = uint(id)
	result := repo.db.Model(model).Where("user_id = ?", userId).Updates(dataReq)
	if result.RowsAffected == 0{
		return errors.New("no row affected")
	}
	if result != nil{
		return result.Error
	}
	
	return  nil

}

func (repo *mysqlProductRepository)SelectDataByUserID(id_user int) (response []products.Core, err error){
	var dataProduct []Product
	result := repo.db.Where("user_id = ?", id_user).Find(&dataProduct)
	if result.Error != nil{
		return []products.Core{}, result.Error
	}
	return ToCoreList(dataProduct),result.Error
}

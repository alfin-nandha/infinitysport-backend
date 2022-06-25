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

func (repo *mysqlProductRepository)DeleteDataByID(id int) (error){
	dataProduct := Product{}
	err := repo.db.Delete(&dataProduct,id)
	return err.Error
}

func (repo *mysqlProductRepository)UpdateDataByID(dataReq map[string]interface{}, id int)(error){

	model := Product{}
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

func (repo *mysqlProductRepository)SelectDataByUserID(id_user int) (response []products.Core, err error){
	var dataProduct []Product
	result := repo.db.Where("id_user = ?", id_user).Find(&dataProduct)
	if result.Error != nil{
		return []products.Core{}, result.Error
	}
	return ToCoreList(dataProduct),result.Error
}

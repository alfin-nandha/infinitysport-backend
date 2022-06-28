package data

import (
	"errors"
	"project/e-comerce/features/carts"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	db *gorm.DB
}

func NewCartRepository(conn *gorm.DB) carts.Data {
	return &mysqlCartRepository{
		db: conn,
	}
}

func (repo *mysqlCartRepository) CheckProductInCart(UserId int, IdProduct int) (result bool, err error) {
	var data Cart
	response := repo.db.Where("user_id = ? AND product_id = ?", UserId, IdProduct).First(&data)

	if response.RowsAffected == 0 {
		return false, response.Error
	}
	if response.RowsAffected != 0 {
		result = true
	}
	return result, nil
}

func (repo *mysqlCartRepository) SelectData() ([]carts.Core, error) {
	var data []Cart

	result := repo.db.Preload("Product").Find(&data)
	if result.Error != nil {
		return []carts.Core{}, result.Error
	}

	return ToCoreList(data), nil
}

func (repo *mysqlCartRepository) InsertData(cart carts.Core) (int, error) {
	Cart := fromCore(cart)
	result := repo.db.Create(&Cart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlCartRepository) Update(UserId int, ProductId int, Qty int) (result int, err error) {
  dataCart := Cart{}
  idCart := repo.db.Where(UserId).Find(&dataCart)
  
 	result := repo.db.Model(&Cart{}, idCart).Update("quantity", Qty)
 	
 	if result.RowsAffected != 0 {
 	  return 0, errors.New("failed to query update")
 	}
 	return result, nil
}

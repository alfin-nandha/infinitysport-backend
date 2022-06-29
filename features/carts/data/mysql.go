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

func (repo *mysqlCartRepository) CheckProductInCart(UserId int, IdProduct int) (result bool, idCart, Qty int, err error) {
	var data Cart
	response := repo.db.Where("user_id = ? AND product_id = ?", UserId, IdProduct).First(&data)

	if response.RowsAffected == 0 {
		return false, 0, 0, nil
	}
	return true, int(data.ID), data.Qty, nil
}

func (repo *mysqlCartRepository) SelectData(UserId int) ([]carts.Core, error) {
	var data []Cart

	result := repo.db.Where("user_id = ?", UserId).Preload("Product").Find(&data)
	if result.Error != nil {
		return []carts.Core{}, result.Error
	}

	return ToCoreList(data), nil
}

func (repo *mysqlCartRepository) InsertData(cart carts.Core) (int, error) {
	Cart := fromCore(cart)
	result := repo.db.Create(&Cart)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlCartRepository) Update(cart carts.Core, idCart int) (result int, err error) {
	tx := repo.db.Model(&Cart{}).Where("id = ?", idCart).Update("qty", cart.Qty)

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to query update")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlCartRepository) Destroy(UserId, idCart int) (result int, err error) {
	tx := repo.db.Where(UserId).Delete(&Cart{}, idCart)

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to query delete")
	}
	return int(tx.RowsAffected), nil
}

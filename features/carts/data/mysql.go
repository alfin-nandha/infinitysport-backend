package data

import (
	"errors"
	"project/e-comerce/features/carts"

	"gorm.io/gorm"
)

type mysqlCartRepository struct{
	db *gorm.DB
}

func NewCartRepository(conn *gorm.DB) carts.Data {
	return &mysqlCartRepository{
		db: conn,
	}
}

func (repo *mysqlCartRepository) InsertCart(data carts.Core) (data carts.Core, err error) {
  
}
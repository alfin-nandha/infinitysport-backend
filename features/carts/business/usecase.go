package business

import (
	"project/e-comerce/features/carts"
)

type cartUseCase struct {
	cartData carts.Data
}

func NewCartBusiness(usrData carts.Data) carts.Business {
	return &cartUseCase{
		cartData: usrData,
	}
}

func (uc *cartUseCase) GetAllCart(UserId int) (response []carts.Core, err error) {
	response, err = uc.cartData.SelectData(UserId)
	return response, err
}

// return valuenya bool, idcart, qty, err
// if not data idcart 0
func (uc *cartUseCase) AddCart(data carts.Core) (result int, err error) {
	resultCart, idCart, Qty, _ := uc.cartData.CheckProductInCart(data.UserID, data.ProductID)

	newQty := Qty + 1

	if !resultCart {
		result, _ = uc.cartData.InsertData(data)
	} else {
		UserId := data.UserID
		result, _ = uc.cartData.Update(UserId, idCart, newQty)
	}
	return result, nil
}

func (uc *cartUseCase) UpdateCart(data carts.Core, idCart int) (result int, err error) {
	UserId := data.UserID
	newQty := data.Qty

	result, _ = uc.cartData.Update(UserId, idCart, newQty)
	return result, nil
}

func (uc *cartUseCase) DestroyCart(UserId, idCart int) (result int, err error) {
	result, _ = uc.cartData.Destroy(UserId, idCart)
	return result, nil
}

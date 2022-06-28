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

func (uc *cartUseCase) GetAllCart() (response []carts.Core, err error) {
	response, err = uc.cartData.SelectData()
	return response, err
}

func (uc *cartUseCase) AddCart(data carts.Core) (result int, err error) {
	resultCart, _ := uc.cartData.CheckProductInCart(data.UserID, data.ProductID)

	if resultCart == false {
		result, err = uc.cartData.InsertData(data)
	}
	return result, err
}

func (uc *cartUseCase) UpdateCart(data carts.Core) (result int, err error) {
  resultCart, _ := uc.cartData.CheckProductInCart(data.UserID, data.ProductID)

  if resultCart == true {
		result, err = uc.cartData.Update(data)
	}
 	return result, nil
}

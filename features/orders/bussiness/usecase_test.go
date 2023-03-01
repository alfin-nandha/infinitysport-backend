package bussiness

import (
	"errors"
	Carts "project/e-comerce/features/carts"
	Orders "project/e-comerce/features/orders"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderDataSuccess struct{}

func (mockOrderDataSuccess) SelectCart(cartid []int, userid int) ([]Orders.OrderDetail, []int, int, error) {
	return []Orders.OrderDetail{
		{ID: 1, Price: 10000, Qty: 2, ProductName: "sepatu baru"},
		{ID: 2, Price: 10000, Qty: 2, ProductName: "sepatu lama"},
	}, []int{1, 2}, 20000, nil
}

func (mockOrderDataSuccess) SelectProduct(productid int) (int, error) {
	return 4, nil
}

func (mockOrderDataSuccess) InsertAddress(address Orders.AddressCore) (adID int, err error) {
	return 1, nil
}

func (mockOrderDataSuccess) InsertOrder(Orders.Core, int) (orderID int, err error) {
	return 1, nil
}

func (mockOrderDataSuccess) InsertOrderDetail(orderID int, dataOrderDetail []Orders.OrderDetail) error {
	return nil
}

func (mockOrderDataSuccess) UpdatePriceOrder(orderID int, totalPrice int) error {
	return nil
}

func (mockOrderDataSuccess) InsertPayment(payment Orders.PaymentCore) error {
	return nil
}

func (mockOrderDataSuccess) SelectOrder(orderID int) ([]Orders.Core, error) {
	return []Orders.Core{
		{ID: 1, AddressID: 1, Price: 20000},
		{ID: 2, AddressID: 2, Price: 20000},
	}, nil
}

func (mockOrderDataSuccess) SelectOrderDetailByOrderID(orderID int) ([]Orders.OrderDetail, error) {
	return []Orders.OrderDetail{
		{ID: 1, ProductName: "sepatu baru", Qty: 2, Price: 10000},
		{ID: 2, ProductName: "sepatu lama", Qty: 2, Price: 10000},
	}, nil
}

func (mockOrderDataSuccess) UpdateStatusOrder(orderID int, userID int, status string) error {
	return nil
}

func (mockOrderDataSuccess) UpdateStockProduct(productID int, qty int) error {
	return nil
}

type mockOrderDataFailed struct{}

func (mockOrderDataFailed) SelectCart(cartid []int, userid int) ([]Orders.OrderDetail, []int, int, error) {
	return []Orders.OrderDetail{}, []int{}, 0, errors.New("failed get cart")
}

func (mockOrderDataFailed) SelectProduct(productid int) (int, error) {
	return 0, errors.New("failed get product")
}

func (mockOrderDataFailed) InsertAddress(address Orders.AddressCore) (adID int, err error) {
	return 0, errors.New("failed add address")
}

func (mockOrderDataFailed) InsertOrder(Orders.Core, int) (orderID int, err error) {
	return 0, errors.New("failed insert order")
}

func (mockOrderDataFailed) InsertOrderDetail(orderID int, dataOrderDetail []Orders.OrderDetail) error {
	return errors.New("failed insert order detail")
}

func (mockOrderDataFailed) UpdatePriceOrder(orderID int, totalPrice int) error {
	return errors.New("failed update price")
}

func (mockOrderDataFailed) InsertPayment(payment Orders.PaymentCore) error {
	return errors.New("failed add payment")
}

func (mockOrderDataFailed) SelectOrder(orderID int) ([]Orders.Core, error) {
	return []Orders.Core{}, errors.New("failed get order")
}

func (mockOrderDataFailed) SelectOrderDetailByOrderID(orderID int) ([]Orders.OrderDetail, error) {
	return []Orders.OrderDetail{}, errors.New("failed get order detail")
}

func (mockOrderDataFailed) UpdateStatusOrder(orderID int, userID int, status string) error {
	return errors.New("failed update status order")
}

func (mockOrderDataFailed) UpdateStockProduct(productID int, qty int) error {
	return errors.New("failed update stock product")
}

type mockCartDataSuccess struct{}

func (mockCartDataSuccess) DestroyAll(userid int, cartid []int) error {
	return nil
}

func (mockCartDataSuccess) CheckProductInCart(UserId int, IdProduct int) (bool, int, int, error) {
	return true, 1, 1, nil
}

func (mockCartDataSuccess) Destroy(userid, cartid int) (int, error) {
	return 1, nil
}

func (mockCartDataSuccess) SelectData(UserId int) (data []Carts.Core, err error) {
	return []Carts.Core{
		{ID: 1, ProductID: 1, Qty: 2},
	}, nil
}
func (mockCartDataSuccess) Update(UserId, idCart, Qty int) (row int, err error) {
	return 1, nil
}

func (mockCartDataSuccess) InsertData(data Carts.Core) (row int, err error) {
	return 1, nil
}

type mockCartDataFailed struct{}

func (mockCartDataFailed) DestroyAll(userId int, cartID []int) error {
	return errors.New("failed delete cart")
}

func (mockCartDataFailed) CheckProductInCart(UserId int, IdProduct int) (bool, int, int, error) {
	return false, 0, 0, errors.New("failed delete cart")
}

func (mockCartDataFailed) Destroy(userid, cartid int) (int, error) {
	return 0, errors.New("failed delete cart")
}

func (mockCartDataFailed) SelectData(UserId int) (data []Carts.Core, err error) {
	return []Carts.Core{}, errors.New("failed delete cart")
}
func (mockCartDataFailed) Update(UserId, idCart, Qty int) (row int, err error) {
	return 0, errors.New("failed delete cart")
}

func (mockCartDataFailed) InsertData(data Carts.Core) (row int, err error) {
	return 0, errors.New("failed delete cart")
}

func TestAddOrder(t *testing.T) {
	t.Run("Test Add Order Success", func(t *testing.T) {
		core := Orders.Core{
			Address: Orders.AddressCore{
				Receiver: "alfin",
				Phone:    "082",
				Address:  "malang",
			},
		}
		cartid := []int{1, 2}
		userid := 1
		orderBussiness := NewOrderBusiness(mockOrderDataSuccess{}, mockCartDataSuccess{})
		err := orderBussiness.AddOrder(core, cartid, userid)
		assert.Nil(t, err)
	})

	t.Run("Test Add Order Failed", func(t *testing.T) {
		core := Orders.Core{}
		cartid := []int{}
		userid := 0
		orderBussiness := NewOrderBusiness(mockOrderDataFailed{}, mockCartDataFailed{})
		err := orderBussiness.AddOrder(core, cartid, userid)
		assert.NotNil(t, err)
	})
}

func TestGetOrder(t *testing.T) {
	t.Run("Test Select Order Success", func(t *testing.T) {
		orderid := 1
		orderBussiness := NewOrderBusiness(mockOrderDataSuccess{}, mockCartDataSuccess{})
		result, err := orderBussiness.GetOrder(orderid)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].ID)
	})

	t.Run("Test Select Order Failed", func(t *testing.T) {
		orderid := 0
		orderBussiness := NewOrderBusiness(mockOrderDataFailed{}, mockCartDataFailed{})
		result, err := orderBussiness.GetOrder(orderid)
		assert.NotNil(t, err)
		assert.Equal(t, []Orders.Core{}, result)
	})
}

func TestGetOrderDetail(t *testing.T) {
	t.Run("Test Select Order Success", func(t *testing.T) {
		orderid := 1
		orderBussiness := NewOrderBusiness(mockOrderDataSuccess{}, mockCartDataSuccess{})
		result, err := orderBussiness.GetOrderDetail(orderid)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].ID)
	})

	t.Run("Test Select Order Failed", func(t *testing.T) {
		orderid := 0
		orderBussiness := NewOrderBusiness(mockOrderDataFailed{}, mockCartDataFailed{})
		result, err := orderBussiness.GetOrderDetail(orderid)
		assert.NotNil(t, err)
		assert.Equal(t, []Orders.OrderDetail{}, result)
	})
}

func TestConfirmOrder(t *testing.T) {
	t.Run("Test Select Order Success", func(t *testing.T) {
		payment := Orders.PaymentCore{
			PaymentName: "transfer",
			PaymentCode: "code123",
			NumberCard:  "123",
		}
		orderid := 1
		userid := 1
		orderBussiness := NewOrderBusiness(mockOrderDataSuccess{}, mockCartDataSuccess{})
		err := orderBussiness.ConfirmOrder(payment, orderid, userid)
		assert.Nil(t, err)
	})

	t.Run("Test Select Order Failed", func(t *testing.T) {
		payment := Orders.PaymentCore{}
		orderid := 0
		userid := 0

		orderBussiness := NewOrderBusiness(mockOrderDataFailed{}, mockCartDataFailed{})
		err := orderBussiness.ConfirmOrder(payment, orderid, userid)
		assert.NotNil(t, err)
	})
}

func TestCancelmOrder(t *testing.T) {
	t.Run("Test Select Order Success", func(t *testing.T) {

		orderid := 1
		userid := 1
		orderBussiness := NewOrderBusiness(mockOrderDataSuccess{}, mockCartDataSuccess{})
		err := orderBussiness.CancelOrder(orderid, userid)
		assert.Nil(t, err)
	})

	t.Run("Test Select Order Failed", func(t *testing.T) {

		orderid := 0
		userid := 0

		orderBussiness := NewOrderBusiness(mockOrderDataFailed{}, mockCartDataFailed{})
		err := orderBussiness.CancelOrder(orderid, userid)
		assert.NotNil(t, err)
	})
}

package bussiness

import (
	"fmt"
	"project/e-comerce/features/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockProductDataSucces struct{}

func (mock mockProductDataSucces) SelectData() (data []products.Core, err error) {
	return []products.Core{
		{ID: 1, Name: "sepatu baru", ProductDetail: "ini sepatu baru", Price: 10000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1},
		{ID: 1, Name: "sepatu lama", ProductDetail: "ini sepatu lama", Price: 1000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1},
	}, nil
}

func (mock mockProductDataSucces) SelectDataByID(id int) (data products.Core, err error) {
	return products.Core{ID: 1, Name: "sepatu baru", ProductDetail: "ini sepatu baru", Price: 10000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1}, nil
}

func (mock mockProductDataSucces) InsertData(data products.Core) (err error) {
	return nil
}

func (mock mockProductDataSucces) DeleteDataByID(id int, userid int) (err error) {
	return nil
}

func (mock mockProductDataSucces) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error) {
	return nil
}

func (mock mockProductDataSucces) SelectDataByUserID(userId int) (data []products.Core, err error) {
	return []products.Core{
		{ID: 1, Name: "sepatu baru", ProductDetail: "ini sepatu baru", Price: 10000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1},
		{ID: 1, Name: "sepatu lama", ProductDetail: "ini sepatu lama", Price: 1000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1},
	}, nil
}

//mock data failed case
type mockProductDataFailed struct{}

func (mock mockProductDataFailed) SelectData() (data []products.Core, err error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockProductDataFailed) SelectDataByID(id int) (data products.Core, err error) {
	return data, fmt.Errorf("Failed to select data")
}

func (mock mockProductDataFailed) InsertData(data products.Core) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockProductDataFailed) DeleteDataByID(id int, userid int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockProductDataFailed) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockProductDataFailed) SelectDataByUserID(userId int) (data []products.Core, err error) {
	return nil, fmt.Errorf("failed to select data ")
}

func TestGetAllProduct(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		productBusiness := NewProductBusiness(mockProductDataSucces{})
		result, err := productBusiness.GetAllProduct()
		assert.Nil(t, err)
		assert.Equal(t, "sepatu baru", result[0].Name)
	})

	t.Run("Test Get All Data Failed", func(t *testing.T) {

		productBusiness := NewProductBusiness(mockProductDataFailed{})
		result, err := productBusiness.GetAllProduct()
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetProductByID(t *testing.T) {
	t.Run("Test Get Product Data By ID Success", func(t *testing.T) {
		id := 1
		productBusiness := NewProductBusiness(mockProductDataSucces{})
		result, err := productBusiness.GetProductByID(id)
		assert.Nil(t, err)
		assert.Equal(t, "sepatu baru", result.Name)
	})

	t.Run("Test Get Product Data By ID Failed", func(t *testing.T) {
		id := 3
		productBusiness := NewProductBusiness(mockProductDataFailed{})
		result, err := productBusiness.GetProductByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name)
	})
}

func TestInsertProduct(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		productBusiness := NewProductBusiness(mockProductDataSucces{})
		newProduct := products.Core{
			Name: "sepatu baru", ProductDetail: "ini sepatu baru", Price: 10000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1,
		}
		err := productBusiness.InsertProduct(newProduct)
		assert.Nil(t, err)
	})

	t.Run("Test Insert Data Failed", func(t *testing.T) {
		productBusiness := NewProductBusiness(mockProductDataFailed{})
		newProduct := products.Core{
			Name: "alta",
		}
		err := productBusiness.InsertProduct(newProduct)
		assert.NotNil(t, err)
	})

}

func TestGetProductByUserID(t *testing.T) {
	t.Run("Test Get Product Data By ID User Success", func(t *testing.T) {
		id := 1
		productBusiness := NewProductBusiness(mockProductDataSucces{})
		result, err := productBusiness.GetProductByUserID(id)
		assert.Nil(t, err)
		assert.Equal(t, "sepatu baru", result[0].Name)
	})

	t.Run("Test Get Data By ID User Failed", func(t *testing.T) {
		id := 3
		productBusiness := NewProductBusiness(mockProductDataFailed{})
		result, err := productBusiness.GetProductByUserID(id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data", func(t *testing.T) {
		id := 1
		userid := 1
		productBusiness := NewProductBusiness(mockProductDataSucces{})
		err := productBusiness.DeleteProductByID(id, userid)
		assert.Nil(t, err)

	})
	t.Run("Test Delete Data Failed", func(t *testing.T) {
		id := 3
		userid := 1
		productBusiness := NewProductBusiness(mockProductDataFailed{})
		err := productBusiness.DeleteProductByID(id, userid)
		assert.NotNil(t, err)

	})
}

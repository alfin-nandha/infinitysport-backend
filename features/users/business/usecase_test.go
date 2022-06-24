package business

import (
	"fmt"
	"project/e-comerce/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockUserData struct{}

func (mock mockUserData) SelectData(param string) (data []users.Core, err error) {
	return []users.Core{
		{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"},
	}, nil
}
func (mock mockUserData) SelectDataById(id int) (data users.Core, err error) {
	return users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}, nil
}


func (mock mockUserData) InsertData(data users.Core) (err error) {
	return nil
}

func (mock mockUserData) DeleteData(id int) (err error) {
	return nil
}

func (mock mockUserData) UpdateData(data map[string]interface{},id int) (err error) {
	return nil
}


//mock data failed case
type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) SelectDataById(id int) (data users.Core, err error) {
	return data, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) InsertData(data users.Core) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockUserDataFailed) DeleteData(id int) (err error) {
	return nil
}

func (mock mockUserDataFailed) UpdateData(data map[string]interface{},id int) (err error) {
	return nil
}



func TestGetAllData(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetAllData("")
		assert.Nil(t, err)
		assert.Equal(t, "alta", result[0].Name)
	})

	t.Run("Test Get All Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetAllData("")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.Nil(t, err)
		//assert.Equal(t, 1, result)
	})

	t.Run("Test Insert Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		//assert.Equal(t, 0, result)
	})

	t.Run("Test Insert Data Failed When Email Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		//assert.Equal(t, -1, result)
	})

	t.Run("Test Insert Data Failed When Password Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:  "alta",
			Email: "alta@mail.id",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		//assert.Equal(t, -1, result)
	})
}
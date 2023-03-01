package business

import (
	"errors"
	Auth "project/e-comerce/features/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAuthDataSucsess struct{}

func (mock mockAuthDataSucsess) FindUser(email string) (resp Auth.Core, err error) {
	return Auth.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "$2a$10$2iESuHQMd8NM2q245RP40esJCpID8o/SyfLiNI2.mHGoGEojEoUe."}, nil
}

type mockAuthDataFailed struct{}

func (mock mockAuthDataFailed) FindUser(email string) (resp Auth.Core, err error) {
	return Auth.Core{}, errors.New("user not found")
}

func TestLogin(t *testing.T) {
	t.Run("Test Login Success", func(t *testing.T) {
		authBusiness := NewAuthBusiness(mockAuthDataSucsess{})
		newUser := Auth.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "123",
		}
		resultToken, resultName, err := authBusiness.Login(newUser)
		assert.Nil(t, err)
		assert.NotNil(t, resultToken)
		assert.Equal(t, "alta", resultName)
	})

	t.Run("Test Login email not found", func(t *testing.T) {
		authBusiness := NewAuthBusiness(mockAuthDataFailed{})
		newUser := Auth.Core{
			Name:     "alta",
			Email:    "abc@mail.id",
			Password: "qwerty",
		}
		resultToken, resultName, err := authBusiness.Login(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, "", resultToken)
		assert.Equal(t, "", resultName)
	})

	t.Run("Test Login Wrong Pass", func(t *testing.T) {
		authBusiness := NewAuthBusiness(mockAuthDataFailed{})
		newUser := Auth.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		resultToken, resultName, err := authBusiness.Login(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, "", resultToken)
		assert.Equal(t, "", resultName)
	})
}

package data

import (
	"project/e-comerce/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role_id  int
	Name     string
	Email    string
	Password string
}

//DTO

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core users.Core) User {
	return User{
		Name:     core.Name,
		Role_id:  core.Role.ID,
		Email:    core.Email,
		Password: core.Password,
	}
}

func toCore(data User) users.Core {
	return data.toCore()
}

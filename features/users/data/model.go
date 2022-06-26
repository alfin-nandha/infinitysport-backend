package data

import (
	"project/e-comerce/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
<<<<<<< HEAD
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
=======
	Role_id  int
	Name     string
	Email    string
	Password string
>>>>>>> 5fbbfb6f86faeeae913389ddbb41649a7c348204
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
<<<<<<< HEAD
	
=======
>>>>>>> 5fbbfb6f86faeeae913389ddbb41649a7c348204
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
		Name: core.Name,
		Email: core.Email,
		Password: core.Password,
	
	}
}

func toCore(data User) users.Core {
	return data.toCore()
}

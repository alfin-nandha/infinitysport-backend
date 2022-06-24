package data

import (
	"project/e-comerce/features/users"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Role_id int `json:"role_id" form:"role_id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role Role
}

type Role struct{
	ID int 
	Role_name string
}

//DTO

func (data *User) toCore() users.Core{
	return users.Core{
		ID: int(data.ID),
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Role: users.Role{
			ID: data.Role.ID,
			Role_name: data.Role.Role_name,
		},
	}
	
}

func toCoreList(data []User)[]users.Core{
	result :=[]users.Core{}
	for key := range data{
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core users.Core) User{
	return User{
		Name: core.Name,
		Role_id: core.Role.ID,
		Email: core.Email,
		Password: core.Password,
		Role: Role{
			ID: core.Role.ID,
			Role_name: core.Role.Role_name,
		},
	}
}

func toCore(data User)users.Core{

	return data.toCore()
}
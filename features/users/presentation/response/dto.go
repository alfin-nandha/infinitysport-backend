package response

import (
	"project/e-comerce/features/users"
	"time"
)

type user struct{
	ID        int      `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	Role 	  role    `json:"role" form:"role"`
}

type role struct{
	Role string
}

func FromCore(data users.Core) user{
	return user{
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		CreatedAt: data.CreatedAt,
		Role: role{
			Role: data.Role.Role_name,
		},
	}
}

func FromCoreList(data []users.Core) []user{
	result := []user{}
	for key := range data{
		result = append(result, FromCore(data[key]))
	}
	return result
}
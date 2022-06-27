package response

import (
	"project/e-comerce/features/users"
)

type user struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func FromCore(data users.Core) user {
	return user{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

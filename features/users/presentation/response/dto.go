package response

import (
	Users "project/e-comerce/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FromCore(data Users.Core) User {
	return User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []Users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

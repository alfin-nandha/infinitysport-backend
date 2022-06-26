package response

import (
	"project/e-comerce/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
<<<<<<< HEAD

}

func FromCore(data users.Core) user{
	return user{
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		CreatedAt: data.CreatedAt,
=======
}

func FromCore(data users.Core) User {
	return User{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
>>>>>>> 5fbbfb6f86faeeae913389ddbb41649a7c348204
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

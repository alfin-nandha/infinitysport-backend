package auth

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Login(Core) (token string, Name string, err error)
}

type Data interface {
	FindUser(data Core) ([]string, error)
}

package carts

type Core struct {
	ID        int
	ProductID int
	UserID    int
	Price     int
	Qty       int
	Product   Product
	User      User
}

type Product struct {
	ID            int
	Name          string
	ProductName   string
	ProductDetail string
	Stock         int
	Price         int
	Photo         string
	PhotoUrl      string
	UserID        int
	User          User
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Business interface {
  AddCart(data Core) (Core, error)
}

type Data interface {
  InsertCart(data Core) (Core, error)
}

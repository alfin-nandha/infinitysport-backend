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
	AddCart(data Core) (row int, err error)
	GetAllCart() (data []Core, err error)
	UpdateCart(data Core) (row int, err error)
}

type Data interface {
	CheckProductInCart(UserId int, IdProduct int) (bool, error)
	InsertData(data Core) (row int, err error)
	SelectData() (data []Core, err error)
	Update(qty int) (row int, err error)
}

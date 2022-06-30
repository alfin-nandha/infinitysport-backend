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
	GetAllCart(UserId int) (data []Core, err error)
	UpdateCart(data Core, Qty int) (row int, err error)
	DestroyCart(UserId, Id int) (row int, err error)
}

type Data interface {
	CheckProductInCart(UserId int, IdProduct int) (bool, int, int, error)
	InsertData(data Core) (row int, err error)
	SelectData(UserId int) (data []Core, err error)
	Update(UserId, idCart, Qty int) (row int, err error)
	Destroy(UserId, id int) (row int, err error)
	DestroyAll(UserId int, cartId []int) error
}

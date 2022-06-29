package routes

import (
	"project/e-comerce/factory"
	"project/e-comerce/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(middlewares.RemoveTrailingSlash())

	e.Use(middlewares.CorsMiddleware())

	e.POST("/signup", presenter.UserPresenter.Insert)
	e.POST("/login", presenter.AuthPresenter.Login)

	e.GET("/users", presenter.UserPresenter.GetAll, middlewares.JWTMiddleware())
	e.GET("/users/:id", presenter.UserPresenter.GetDataById, middlewares.JWTMiddleware())
	e.PUT("/users/:id", presenter.UserPresenter.Update)
	e.DELETE("/users/:id", presenter.UserPresenter.Delete, middlewares.JWTMiddleware())

	e.GET("/products", presenter.ProductPresenter.GetAll)
	e.GET("/products/:id", presenter.ProductPresenter.GetDataById)
	e.POST("/products", presenter.ProductPresenter.InsertData, middlewares.JWTMiddleware())
	e.PUT("/products/:id", presenter.ProductPresenter.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", presenter.ProductPresenter.DeleteData, middlewares.JWTMiddleware())

	e.GET("/user-products", presenter.ProductPresenter.GetProductByUser, middlewares.JWTMiddleware())

	e.POST("/orders", presenter.OrderPresenter.AddOrder, middlewares.JWTMiddleware())
	e.GET("/orders", presenter.OrderPresenter.GetOrder, middlewares.JWTMiddleware())
	e.GET("/orders/:orderid", presenter.OrderPresenter.GeOrderDetailByOrderID, middlewares.JWTMiddleware())
	e.POST("/orders/:orderid/confirm", presenter.OrderPresenter.ConfirmOrder, middlewares.JWTMiddleware())
	e.PUT("/orders/:orderid/cancel", presenter.OrderPresenter.CancelOrder, middlewares.JWTMiddleware())
	// Cart
	e.POST("/carts", presenter.CartPresenter.AddCart, middlewares.JWTMiddleware())
	e.GET("/carts", presenter.CartPresenter.GetAll, middlewares.JWTMiddleware())
	e.PUT("/carts/:id", presenter.CartPresenter.Update, middlewares.JWTMiddleware())
	e.DELETE("/carts/:id", presenter.CartPresenter.Destroy, middlewares.JWTMiddleware())

	return e
}

//AKIA3JBST3XEWGFLCPYY
//CNWNe7ZuXs9PsJwJxmAnxblCt7gAnO6qnppsVtrJ

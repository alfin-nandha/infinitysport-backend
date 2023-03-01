package routes

import (
	Factory "project/e-comerce/factory"
	Middlewares "project/e-comerce/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter Factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(Middlewares.RemoveTrailingSlash())

	e.Use(Middlewares.CorsMiddleware())

	e.POST("/signup", presenter.UserPresenter.Insert)
	e.POST("/login", presenter.AuthPresenter.Login)

	e.GET("/users", presenter.UserPresenter.GetAll, Middlewares.JWTMiddleware())
	e.GET("/users/details", presenter.UserPresenter.GetDataById, Middlewares.JWTMiddleware())
	e.PUT("/users", presenter.UserPresenter.Update, Middlewares.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.Delete, Middlewares.JWTMiddleware())

	e.GET("/products", presenter.ProductPresenter.GetAll)
	e.GET("/products/:id", presenter.ProductPresenter.GetDataById)
	e.POST("/products", presenter.ProductPresenter.InsertData, Middlewares.JWTMiddleware())
	e.PUT("/products/:id", presenter.ProductPresenter.UpdateData, Middlewares.JWTMiddleware())
	e.DELETE("/products/:id", presenter.ProductPresenter.DeleteData, Middlewares.JWTMiddleware())

	e.GET("/user-products", presenter.ProductPresenter.GetProductByUser, Middlewares.JWTMiddleware())

	e.POST("/orders", presenter.OrderPresenter.AddOrder, Middlewares.JWTMiddleware())
	e.GET("/orders", presenter.OrderPresenter.GetOrder, Middlewares.JWTMiddleware())
	e.GET("/orders/:orderid", presenter.OrderPresenter.GeOrderDetailByOrderID, Middlewares.JWTMiddleware())
	e.POST("/orders/:orderid/confirm", presenter.OrderPresenter.ConfirmOrder, Middlewares.JWTMiddleware())
	e.PUT("/orders/:orderid/cancel", presenter.OrderPresenter.CancelOrder, Middlewares.JWTMiddleware())

	e.POST("/carts", presenter.CartPresenter.AddCart, Middlewares.JWTMiddleware())
	e.GET("/carts", presenter.CartPresenter.GetAll, Middlewares.JWTMiddleware())
	e.PUT("/carts/:id", presenter.CartPresenter.Update, Middlewares.JWTMiddleware())
	e.DELETE("/carts/:id", presenter.CartPresenter.Destroy, Middlewares.JWTMiddleware())

	return e
}

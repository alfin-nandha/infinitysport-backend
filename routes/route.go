package routes

import (
	"project/e-comerce/factory"
	"project/e-comerce/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(middlewares.RemoveTrailingSlash())
	//e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/auth", presenter.AuthPresenter.Login)

	e.GET("/users", presenter.UserPresenter.GetAll, middlewares.JWTMiddleware())
	e.GET("/users/:id", presenter.UserPresenter.GetDataById, middlewares.JWTMiddleware())
	e.POST("/users", presenter.UserPresenter.Insert)
	e.PUT("/users/:id", presenter.UserPresenter.Update)
	e.DELETE("/users/:id", presenter.UserPresenter.Delete, middlewares.JWTMiddleware())

	e.GET("/products", presenter.ProductPresenter.GetAll)
	e.GET("/products/:id", presenter.ProductPresenter.GetDataById)
	e.POST("/products", presenter.ProductPresenter.InsertData, middlewares.JWTMiddleware())
	e.PUT("/products/:id", presenter.ProductPresenter.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", presenter.ProductPresenter.DeleteData, middlewares.JWTMiddleware())
	
	e.GET("/user-products", presenter.ProductPresenter.GetProductByUser, middlewares.JWTMiddleware())

	return e
}

//AKIA3JBST3XEWGFLCPYY
//CNWNe7ZuXs9PsJwJxmAnxblCt7gAnO6qnppsVtrJ
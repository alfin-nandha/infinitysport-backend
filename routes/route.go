package routes

import (
	"project/e-comerce/factory"
	"project/e-comerce/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(middlewares.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

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

	return e
}

//AKIA3JBST3XEWGFLCPYY
//CNWNe7ZuXs9PsJwJxmAnxblCt7gAnO6qnppsVtrJ

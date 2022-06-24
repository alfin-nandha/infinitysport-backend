package main

import (
	"project/e-comerce/factory"
	"project/e-comerce/middlewares"
	"project/e-comerce/routes"

	"project/e-comerce/config"
)

func main(){
	dbConn := config.InitDB()
	presenter := factory.InitFactory(dbConn)
	e :=routes.New(presenter)
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}

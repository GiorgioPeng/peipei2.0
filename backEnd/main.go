package main

import (
	"github.com/kataras/iris/v12"
	"peipei2/router"
)

func main() {

	app := router.NewApp()
	err := app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}

}

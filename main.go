package main

import "github.com/kataras/iris"

func main() {
	app:=iris.New()
	app.Get("/", func(c iris.Context) {
	})
	app.Run(iris.Addr(":8080"))
}

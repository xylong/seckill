package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"seckill/backend/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML("./backend/web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(template)
	app.StaticWeb("/assets", "./backend/web/assets")
	app.OnAnyErrorCode(func(c iris.Context) {
		c.ViewData("msg", c.Values().GetStringDefault("msg", "访问页面出错"))
		c.ViewLayout("")
		c.View("shared/error.html")
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx)
	product.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

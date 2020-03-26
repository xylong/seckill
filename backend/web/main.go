package main

import "github.com/kataras/iris"

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
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

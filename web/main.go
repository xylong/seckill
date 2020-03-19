package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"seckill/datasource"
	"seckill/repositories"
	"seckill/services"
	"seckill/web/controllers"
	"seckill/web/middleware"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// 加载视图模板地址
	app.RegisterView(iris.HTML("./web/views", ".html"))

	mvc.Configure(app.Party("/movies"), movies)

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func movies(app *mvc.Application) {
	app.Router.Use(middleware.BasicAuth)

	repo := repositories.NewMovieRepository(datasource.Movies)
	movieService := services.NewMovieService(repo)
	app.Register(movieService)
	app.Handle(new(controllers.MovieController))
}

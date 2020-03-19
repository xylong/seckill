package controllers

import (
	"errors"
	"github.com/kataras/iris"
	"seckill/datamodels"
	"seckill/services"
)

type MovieController struct {
	Service services.MovieService
}

func (c *MovieController) Get() (results []datamodels.Movie) {
	return c.Service.GetAll()
}

func (c *MovieController) GetBy(id int64) (movie datamodels.Movie, found bool) {
	return c.Service.GetByID(id)
}

func (c *MovieController) PutBy(ctx iris.Context, id int64) (datamodels.Movie, error) {
	// 获取请求内的 poster 和 genre 字段数据
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels.Movie{}, errors.New("failed due form file 'poster' missing")
	}
	// 关闭
	file.Close()

	// 假设这是一个上传文件的 url ...
	poster := info.Filename
	genre := ctx.FormValue("genre")

	return c.Service.UpdatePosterAndGenreByID(id, poster, genre)
}

func (c *MovieController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		// 返回被删除的 movie 的 id
		return iris.Map{"deleted": id}
	}
	//在这里，我们可以看到一个方法函数可以返回两种类型中的任何一种（map 或者 int）,
	// 我们不用指定特定的返回类型。
	return iris.StatusBadRequest
}

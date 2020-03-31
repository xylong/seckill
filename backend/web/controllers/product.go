package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"seckill/services"
)

type ProductController struct {
	c       iris.Context
	service services.ProductService
}

func (p *ProductController) GetAll() mvc.View {
	products, _ := p.service.GetAllProduct()
	return mvc.View{
		Name: "product/index.html",
		Data: iris.Map{
			"products": products,
		},
	}
}

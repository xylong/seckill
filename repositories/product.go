package repositories

import (
	"seckill/datamodels"
)

type Product interface {
	Insert(product *datamodels.Product) (id int64, err error)
	Delete(id int64) (err error)
	Update(product *datamodels.Product) (err error)
	Get(id int64) (product *datamodels.Product, err error)
	Select() (products []*datamodels.Product, err error)
}

func NewProduct() Product {
	return &ProductManager{}
}

type ProductManager struct {
}

func (p *ProductManager) Insert(product *datamodels.Product) (id int64, err error) {
	sql := "INSERT INTO product(name,number,image,url) VALUE(?,?,?,?)"
	result, err := db.Exec(sql, product.Number, product.Number, product.Image, product.Url)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

func (p *ProductManager) Delete(id int64) (err error) {
	return
}

func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	return
}

func (p *ProductManager) Get(id int64) (product *datamodels.Product, err error) {
	return
}

func (p *ProductManager) Select() (products []*datamodels.Product, err error) {
	return
}

package repositories

import (
	"github.com/jmoiron/sqlx"
	"seckill/datamodels"
)

type Product interface {
	Insert(product *datamodels.Product) (id int64, err error)
	Delete(id int64) bool
	Update(product *datamodels.Product) (err error)
	Get(id int64) (product *datamodels.Product, err error)
	Select(ids []int64) (products []*datamodels.Product, err error)
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

func (p *ProductManager) Delete(id int64) bool {
	sql := "DELETE FROM product WHERE id=?"
	res, err := db.Exec(sql, id)
	if err != nil {
		return false
	}
	n, err := res.RowsAffected()
	if err != nil || n == 0 {
		return false
	}
	return true
}

func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	sql := "UPDATE PRODUCT SET name=?,number=?,image=?,url=? WHERE id=?"
	res, err := db.Exec(sql, product.Name, product.Number, product.Url, product.Url, product.ID)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil || n == 0 {
		return err
	}
	return nil
}

func (p *ProductManager) Get(id int64) (product *datamodels.Product, err error) {
	product = &datamodels.Product{}
	sql := "SELECT id,name,number,image,url FROM PRODUCT WHERE id=?"
	err = db.Get(product, sql, id)
	return
}

func (p *ProductManager) Select(ids []int64) (products []*datamodels.Product, err error) {
	sql, args, err := sqlx.In("SELECT id,name,number,image,url FROM PRODUCT WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	err = db.Select(&products, sql, args...)
	if err != nil {
		return
	}
	return
}

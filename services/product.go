package services

import (
	"seckill/datamodels"
	"seckill/repositories"
)

type ProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
}

func NewProductService() ProductService {
	return &Ps{repositories.NewProduct()}
}

type Ps struct {
	Product repositories.Product
}

func (p *Ps) GetProductByID(id int64) (*datamodels.Product, error) {
	return p.Product.Get(id)
}

func (p *Ps) GetAllProduct() ([]*datamodels.Product, error) {
	return p.Product.Select([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func (p *Ps) DeleteProductByID(id int64) bool {
	return p.Product.Delete(id)
}

func (p *Ps) UpdateProduct(product *datamodels.Product) error {
	return p.Product.Update(product)
}

func (p *Ps) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.Product.Insert(product)
}

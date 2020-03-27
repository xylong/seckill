package repositories

import (
	"seckill/datamodels"
	"testing"
)

func TestProductManager_Insert(t *testing.T) {
	product := &datamodels.Product{
		Name:   "MacBook",
		Number: 10,
		Image:  "https://img12.360buyimg.com/n7/jfs/t1/86855/3/15583/138431/5e72575cE318939c3/8bca468b120af220.jpg",
		Url:    "https://item.jd.com/100011979232.html",
	}
	p := NewProduct()
	id, err := p.Insert(product)
	if err != nil {
		panic(err)
	}
	t.Logf("id:%d\n", id)
}

func TestProductManager_Delete(t *testing.T) {
	p := NewProduct()
	ok := p.Delete(6)
	t.Logf("ok:%v\n", ok)
}

func TestProductManager_Update(t *testing.T) {
	product := &datamodels.Product{
		ID:     7,
		Name:   "apple",
		Number: 100,
		Image:  "https://img12.360buyimg.com/n7/jfs/t1/86855/3/15583/138431/5e72575cE318939c3/8bca468b120af220.jpg",
		Url:    "https://item.jd.com/100011979232.html",
	}
	p := NewProduct()
	err := p.Update(product)
	if err != nil {
		panic(err)
	}
}

func TestProductManager_Get(t *testing.T) {
	p := NewProduct()
	product, err := p.Get(8)
	if err != nil {
		panic(err)
	}
	t.Logf("product:%v\n", product)
}

func TestProductManager_Select(t *testing.T) {
	p := NewProduct()
	ids := []int64{3, 4, 5}
	products, err := p.Select(ids)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		t.Logf("product:%#v\n", product)
	}
}

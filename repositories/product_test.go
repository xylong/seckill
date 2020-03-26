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

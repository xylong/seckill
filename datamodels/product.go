package datamodels

type Product struct {
	ID     int64  `json:"id" sql:"id"`
	Name   string `json:"name" sql:"name"`
	Number int64  `json:"number" sql:"number"`
	Image  string `json:"image" sql:"image"`
	Url    string `json:"url" sql:"url"`
}

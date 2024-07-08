package model

type Sale struct {
	NameProduct string  `json:"name_prod" validate:"required"`
	Company     string  `json:"company" validate:"required,email"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Count       int     `json:"count" validate:"required,gt=0.0"`
}

type InfoSales struct {
	NameProduct string
	Company     string
	Price       float64
	Count       int
	Total       float64
}

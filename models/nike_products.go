package models


type NikeProduct struct {
	ProductID 				string	`json:"product_id"`
	Name          			string 	`json:"name"`
	Color         			string 	`json:"color"`
	StyleID       			string 	`json:"style_id"`
	Price         			float64  `json:"price"`
	CurrencyCode  			string 	`json:"currency_code"`
	ProductImage 			string	`json:"product_image"`
	Sizes 					[]string	`json:"sizes"`
}
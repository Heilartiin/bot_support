package models

import "time"

type NapProduct struct {
	ID 				int			`json:"id" db:"id" csv:"id"`
	Name 			string	    `json:"name" db:"name" csv:"name"`
	SizeFit 		string		`json:"size_fit" db:"size_fit" csv:"size_fit"`
	Price 			float64		`json:"price" db:"price" csv:"price"`
	Currency 		string		`json:"currency" db:"currency" csv:"currency"`
	ImageURL 		string		`json:"image_url" db:"image_url" csv:"image_url"`
	ColorID 		int			`json:"color_id" db:"color_id" csv:"color_id"`
	ProductURL 		string		`json:"product_url" db:"product_url" csv:"product_url"`
	BrandName 		string		`json:"brand_name" db:"brand_name" csv:"brand_name"`
	AnalyticsKey	string		`json:"analytics_key" db:"analytics_key" csv:"analytics_key"`
	Invisible 		bool		`json:"invisible" db:"invisible" csv:"invisible"`
	CreatedAT 		time.Time	`json:"created_at" db:"created_at" csv:"created_at"`
	UpdatedAT 		time.Time	`json:"updated_at" db:"updated_at" csv:"-"`
	DeletedAT 		time.Time	`json:"deleted_at" db:"deleted_at" csv:"-"`
}

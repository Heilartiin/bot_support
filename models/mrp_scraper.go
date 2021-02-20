package models

import "time"

// id
//name
//label
//designer_name
//variant_part_number
//part_number
//size
//image_url
//created_at
//updated_at
//deleted_at

type ScraperProduct struct {
	ID 					int		`json:"id" db:"id"`
	Name 				string	`json:"name" db:"name"`
	Label 				string	`json:"label" db:"label"`
	DesignerName		string 	`json:"designer_name" db:"designer_name"`
	ValiantPartNumber	string	`json:"valiant_part_number" db:"valiant_part_number"`
	PartNumber 			string	`json:"part_number" db:"part_number"`
	Size 				string	`json:"size" db:"size"`
	ImageUrl 			string	`json:"image_url" db:"image_url"`
	CreatedAt 			time.Time 	`json:"created_at" db:"created_at"`
	UpdatedAt 			time.Time 	`json:"updated_at" db:"updated_at"`
	DeletedAt 			time.Time 	`json:"deleted_at" db:"deleted_at"`
}

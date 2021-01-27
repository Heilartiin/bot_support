package models

import (
	"encoding/json"
)

type Product struct {
	ID 			int 		`json:"id" db:"id"`
	TaskID 		int 		`json:"task_id" db:"task_id"`
	StoreID		string		`json:"store_id" db:"store_id"`
	WishListID 	string		`json:"wish_list_id" db:"wish_list_id"`
	AccessKey	string 		`json:"access_key" db:"access_key"`
	Pid 		string		`json:"pid" db:"pid"`
	Name 		string 		`json:"name" db:"name"`
	Image 		string 		`json:"image" db:"image"`
	Price 		int			`json:"price" db:"price"`
	Symbol 		string		`json:"symbol" db:"symbol"`
	StockLevel 	int 		`json:"stock_level" db:"stock_level"`
	Sizes 		[]*Size 	`json:"sizes"`
	Message 	string		`json:"message"`
	//CreatedAt 	time.Time 	`json:"created_at" db:"created_at"`
	//UpdatedAt 	time.Time 	`json:"updated_at" db:"updated_at"`
	//DeletedAt 	time.Time 	`json:"deleted_at" db:"deleted_at"`
}

type Size struct {
	ID 				int 		`json:"id" db:"id"`
	TaskID 			int 		`json:"task_id" db:"task_id"`
	Pid 			string 		`json:"pid" db:"pid"`
	Name 			string		`json:"name"`
	Buyable			bool		`json:"buyable" db:"buyable"`
	Image 			string  	`json:"image"`
	Price 			int			`json:"price"`
	Symbol 			string		`json:"symbol"`
	EIP				bool		`json:"eip" db:"eip"`
	Visible 		bool		`json:"visible" db:"visible"`
	Displayable		bool 		`json:"displayable" db:"displayable"`
	Banned			bool		`json:"banned" db:"banned"`
	SizeName 		string		`json:"size_name" db:"size_name"`
	SizeChart		string		`json:"size_chart"`
	StockLevel		int			`json:"stock_level" db:"stock_level"`
	PartNumber		string  	`json:"part_number" db:"part_number"`
	//CreatedAt 		time.Time 	`json:"created_at" db:"created_at"`
	//UpdatedAt 		time.Time 	`json:"updated_at" db:"updated_at"`
	//DeletedAt 		time.Time 	`json:"deleted_at" db:"deleted_at"`
}


type RawProduct struct {
	Product 	*Product 		`db:"product"`
	Sizes 		json.RawMessage `db:"sizes"`
}


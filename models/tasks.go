package models

import "time"

type Task struct {
	ID 			int			`json:"id" db:"id"`
	Pid 		string 		`json:"pid" db:"pid"`
	Title 		string		`json:"title" db:"title"`
	StoreID		string		`json:"store_id" db:"store_id"`
	WishListID	string		`json:"wish_list_id" db:"wish_list_id"`
	AccessKey 	string		`json:"access_key" db:"access_key"`
	Active		bool 		`json:"active" db:"active"`
	Alarm		bool		`json:"alarm" db:"alarm"`
	TimeSleep   int			`json:"time_sleep" db:"time_sleep"`
	CreatedAt 	time.Time 	`json:"-" db:"created_at"`
	UpdatedAt 	time.Time 	`json:"-" db:"updated_at"`
	DeletedAt 	time.Time 	`json:"-" db:"deleted_at"`
}
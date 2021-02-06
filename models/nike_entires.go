package models

import "time"

type NikeEntry struct {
	ID 			int			`json:"id" db:"id"`
	Launch  	string		`json:"launch" csv:"launch" db:"launch"`
	Username 	string		`json:"username" csv:"username" db:"username"`
	Password 	string		`json:"password" csv:"password" db:"password"`
	EntryTime 	time.Time	`json:"entry_time" csv:"entry_time" db:"entry_time"`
	Status 		string		`json:"status" csv:"status" db:"status"`
	// entered = 0 - not entered, 1 - entered
	Entered 	int			`json:"entered" csv:"entered" db:"entered"`
	StyleID 	string		`json:"style_id" csv:"style_id" db:"style_id"`
	CreatedAt   time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at" db:"updated_at"`
	DeletedAt	time.Time	`json:"deleted_at" db:"deleted_at"`
}

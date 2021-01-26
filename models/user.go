package models

import "time"

type User struct {
	ID 					int			`json:"id" db:"id"`
	MemberID 			string		`json:"member_id" db:"member_id"`
	WishList 			string		`json:"wish_list" db:"wishlist"`
	Token 				string		`json:"token" db:"token"`
	PrivateChannel 		string		`json:"private_channel" db:"private_channel"`
	CreatedAt 			time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt  			time.Time	`json:"updated_at" db:"updated_at"`
	DeletedAt 			time.Time	`json:"deleted_at" db:"deleted_at"`
}

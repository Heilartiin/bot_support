package models

import "time"

var PriorityName = map[string]int{
	"general": 1,
	"arthur":  2,
	"dmitriy": 3,
	"other":   4,
}

var Priority = map[int]string{
	1: "general",
	2: "arthur",
	3: "dmitriy",
	4: "other",
}


type NikeAccount struct {
	ID 			int			`json:"id" db:"id" csv:"id"`
	Login 		string		`json:"login" db:"login" csv:"login"`
	Password 	string		`json:"password" db:"password" csv:"password"`
	Priority	int			`json:"priority" db:"priority" csv:"priority"`
	CreatedAt   time.Time	`json:"created_at" db:"created_at" csv:"-"`
	UpdatedAt	time.Time	`json:"updated_at" db:"updated_at" csv:"-"`
	DeletedAt	time.Time	`json:"deleted_at" db:"deleted_at" csv:"-"`
}
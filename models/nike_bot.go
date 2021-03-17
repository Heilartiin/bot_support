package models

import "time"

type BotConfig struct {
	DelayMin 			int         			`json:"delay_min"`
	DelayMax 			int         			`json:"delay_max"`
	SubmitRetry			int      				`json:"submit_retry"`
	ReleaseTime     	time.Time   			`json:"release_time"`
	PrepareSessionTime	time.Time 				`json:"prepare_session_time"`
	EntryTime			time.Time   			`json:"entry_time"`
	Tasks   		 	[]*NikeBotTask          `json:"tasks"`
}


type NikeBotTask struct {
	Login 				string		`json:"login" db:"login"`
	Password 			string		`json:"password" db:"password"`
	SizeName 			string		`json:"size_name"`
	ProductID 			string		`json:"product_id"`
	LaunchID 			string		`json:"launch_id"`
}


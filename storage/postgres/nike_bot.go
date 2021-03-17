package postgres

import "github.com/Heilartin/bot_support/models"

func (db *DB) GetAccountAllActive() ([]*models.NikeBotTask, error) {
	var res []*models.NikeBotTask
	err := db.DB.Select(&res, "SELECT login, password FROM nike_bot_accounts WHERE active = true")
	if err != nil {
		return nil, err
	}
	return res, nil
}


package postgres

import "github.com/Heilartin/bot_support/models"

func (db *DB) GetNapProductInvisible() ([]*models.NapProduct, error)  {
	query := `SELECT * FROM nap_legacy_products WHERE invisible=true`
	var res []*models.NapProduct
	err := db.DB.Select(&res, query)
	if err != nil {
		db.Logger.Error(err)
		return nil, err
	}
	return res, nil
}

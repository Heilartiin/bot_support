package postgres

import "github.com/Heilartin/bot_support/models"

func (db *DB) GetScrapersProductByPid(pid string) ([]*models.ScraperProduct, error) {
	var res []*models.ScraperProduct
	query := `SELECT * FROM mrp_scraper WHERE variant_part_number = $1;`
	err := db.DB.Select(&res, query, pid)
	if err != nil {
		db.Logger.Error(err)
		return nil, err
	}
	return res, nil
}


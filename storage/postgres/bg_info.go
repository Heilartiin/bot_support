package postgres

import (
	"database/sql"
	"encoding/json"
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
)

func (db *DB) SelectBgInfo() (*models.BGInfo, error) {
	var res []*models.RawBgInfo
	query := "SELECT * FROM bg_info"
	err := db.DB.Select(&res, query)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	if len(res) == 0 {
		db.Logger.Error(errors.WithStack(sql.ErrNoRows))
		return nil, errors.WithStack(sql.ErrNoRows)
	}
	r, err := db.rawProcessingBGInfo(res[0])
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return r, nil
}



func (db *DB) rawProcessingBGInfo(m *models.RawBgInfo) (*models.BGInfo, error) {
	var proxies []string

	err := json.Unmarshal(m.Proxies, &proxies)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	res := models.BGInfo{
		Proxies:      proxies,
	}

	return &res, nil
}

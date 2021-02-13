package postgres

import (
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
	"time"
)

func (db *DB) CreateNikeEntry(i *models.NikeEntry) (int, error) {
	nikeEntered := 0
	query := `INSERT INTO nike_entries (launch, username, password, entry_time, status, entered,
 			style_id, created_at, updated_at, deleted_at) VALUES
			( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;`
	err := db.DB.QueryRow(query, i.Launch, i.Username, i.Password, i.EntryTime,
		i.Status, i.Entered, i.StyleID, time.Now(), time.Now(), time.Time{}).Scan(&nikeEntered)
	if err != nil {
		db.Logger.Error(err, i)
		return nikeEntered, err
	}
	return nikeEntered, nil
}


func (db *DB) StorageNikeEntries(in []*models.NikeEntry) (int, error) {
	query := `INSERT INTO nike_entries (launch, username, password, entry_time, status, entered,
 			style_id, created_at, updated_at, deleted_at) VALUES
			( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	tr, err := db.DB.Begin()
	if err != nil {
		db.Logger.Error(err)
		return 0, err
	}
	var NR int64
	for _, i := range in {
		res, err := tr.Exec(query, i.Launch, i.Username, i.Password, i.EntryTime,
			i.Status, i.Entered, i.StyleID, time.Now(), time.Now(), time.Time{})
		if err != nil {
			if err.Error() =="pq: duplicate key value violates unique constraint \"nike_entries_username_launch_key\"" {
				continue
			}
			db.Logger.Error(err)
			tr.Rollback()
			return 0, err
		}
		newRows, err := res.RowsAffected()
		if err != nil {
			db.Logger.Error(err)
			tr.Rollback()
			return 0, err
		}
		NR += newRows
	}
	err = tr.Commit()
	if err != nil {
		db.Logger.Error(err)
		return 0, err
	}
	db.Logger.Info (fmt.Sprintf("%d new rows inserted", NR))
	return int(NR), nil
}

func (db *DB) GetNikeEntriesSortByTime(launch string) ([]*models.NikeEntry, error)  {
	var res []*models.NikeEntry
	query := "SELECT * FROM nike_entries WHERE launch=$1 AND entry_time != '0001-01-01' ORDER BY entry_time;"
	err := db.DB.Select(&res, query, launch)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}
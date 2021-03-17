package postgres

import (
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
	"time"
)

func (db *DB) CreateNikeAccount(i *models.NikeAccount) (int, error) {
	nikeEntered := 0
	query := `INSERT INTO nike_accounts (
	login, password, priority, created_at, updated_at, deleted_at) VALUES
			( $1, $2, $3, $4, $5, $6) RETURNING id;`
	err := db.DB.QueryRow(
		query, i.Login, i.Password, i.Priority, time.Now(),
		time.Now(), time.Time{}).Scan(&nikeEntered)
	if err != nil {
		return nikeEntered, err
	}
	return nikeEntered, nil
}


func (db *DB) StorageNikeAccount(in []*models.NikeAccount) (int, error) {
	query := `INSERT INTO nike_accounts (
	login, password, priority, created_at, updated_at, deleted_at) VALUES
			( $1, $2, $3, $4, $5, $6) RETURNING id;`
	tr, err := db.DB.Begin()
	if err != nil {
		db.Logger.Error(err)
		return 0, err
	}
	var NR int64
	for _, i := range in {
		res, err := tr.Exec(query,i.Login, i.Password, i.Priority, time.Now(),
			time.Now(), time.Time{})
		if err != nil {
			if err.Error() =="pq: duplicate key value violates unique constraint \"nike_accounts_login_key\"" {
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

func (db *DB) GetAllAccounts() ([]*models.NikeAccount, error) {
	var res []*models.NikeAccount
	query := "SELECT * FROM nike_accounts ORDER BY id;"
	err := db.DB.Select(&res, query)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetAccountsByPriority(priorityID int) ([]*models.NikeAccount, error) {
	var res []*models.NikeAccount
	query := "SELECT * FROM nike_accounts WHERE priority = $1;"
	err := db.DB.Select(&res, query, priorityID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

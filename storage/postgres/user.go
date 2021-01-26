package postgres

import (
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateUser(u models.User) (int, error) {
	var UserID int
	query := `INSERT INTO users (member_id, token, wishlist, private_channel,
 				created_at,updated_at, deleted_at) VALUES
			(:member_id, :token, :wishlist, :private_channel, now(), now(), :deleted_at) RETURNING id;`
	rows, err := db.DB.NamedQuery(query, u)
	if err != nil {
		return UserID, errors.WithStack(err)
	}
	if rows.Next() {
		err := rows.Scan(&UserID)
		return UserID, errors.WithStack(err)
	}
	return UserID, nil
}

func (db *DB) GetUserByMemberID(memberID string) (models.User, error) {
	var res models.User
	query := `SELECT * FROM users WHERE member_id = $1 AND users.deleted_at = '0001-01-01' OR users.deleted_at IS NULL;`
	err := db.DB.Get(&res, query, memberID)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}



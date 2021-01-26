package storage

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/storage/postgres"
	"github.com/pkg/errors"
)

type Storage struct {
	Logger *logger.LocalLogger
	DB     *postgres.DB
}

func NewStorage(cfg *config.Config, log *logger.LocalLogger) (*Storage, error) {
	res := &Storage{}
	db, err := postgres.NewPostgres(cfg, log)
	if err != nil {
		return res, errors.WithStack(err)
	}
	bgInfo, err := db.SelectBgInfo()
	if err != nil {
		return res, errors.WithStack(err)
	}
	res.DB = db
	res.Logger = log
	res.DB.BGInfo = bgInfo
	return res, nil
}

package postgres

import (
	"fmt"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	Logger   *logger.LocalLogger
	DB       *sqlx.DB
	DBConfig *config.Config
	BGInfo   *models.BGInfo
}

func NewPostgres(cfg *config.Config, logger *logger.LocalLogger) (*DB, error) {
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBSchema)
	ommitesStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, "[ommited]", cfg.DBName)
	db, err := sqlx.Connect("postgres", str)
	if err != nil {
		logger.Errorf("could not establish connection to ", ommitesStr, err)
		return nil, errors.WithStack(err)
	}
	//db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(50)
	_, err = db.Exec(fmt.Sprintf("set search_path='%s'", cfg.DBSchema))
	if err != nil {
		logger.Errorf("set search_path to ", ommitesStr, err)
		return nil, errors.WithStack(err)
	}

	res := &DB{
		DB:       db,
		Logger:   logger,
		DBConfig: cfg,
	}

	return res, nil
}

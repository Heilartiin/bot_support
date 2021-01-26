package mrporter

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/storage"
	"net/http"
	"time"
)

type MrpClient struct {
	cfg 		*config.MrpPorter
	repo  	 	*storage.Storage
	client 		*http.Client
	log    		*logger.LocalLogger
}

func NewMrpClient(log *logger.LocalLogger, cfg *config.MrpPorter, repo *storage.Storage) *MrpClient  {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	c := MrpClient{
		cfg: cfg,
		repo: repo,
		log: log,
		client: &client,
	}
	return &c
}

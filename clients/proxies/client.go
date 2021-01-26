package proxies

import (
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/storage"
)

type FClient struct {
	log    		*logger.LocalLogger
	repo   		*storage.Storage
	proxies 	[]string
}

func NewFileClient(log *logger.LocalLogger, repo   *storage.Storage) *FClient {
	return &FClient{
		log:    log,
		repo:   repo,
		proxies: repo.DB.BGInfo.Proxies,
	}
}

package opensea

import (
	"github.com/Heilartin/bot_support/clients/ethereum"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"net/http"
	"time"
)

type Client struct {
	httpClient 	*http.Client
	cfg 		*config.OpenSea
	log 		*logger.LocalLogger
	EthClient 	*ethereum.Client
}

func NewClient(log *logger.LocalLogger, cfg *config.Config) (client *Client)  {
	httpClient := http.Client{Timeout: 5 * time.Second}
	return &Client{
		httpClient: &httpClient,
		cfg:        cfg.OpenSea,
		EthClient:  ethereum.NewClient(log, cfg.Infura),
		log:        log,
	}
}

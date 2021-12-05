package opensea

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"net/http"
	"time"
)

type Client struct {
	httpClient 	*http.Client
	cfg 		*config.OpenSea
	log 		*logger.LocalLogger
}

func NewClient(log *logger.LocalLogger, cfg *config.OpenSea) (client *Client)  {
	httpClient := http.Client{Timeout: 5 * time.Second}
	return &Client{
		httpClient: &httpClient,
		cfg:        cfg,
		log:        log,
	}
}

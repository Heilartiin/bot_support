package _cloud

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"net/http"
	"time"
)

type OneCClient struct {
	log    		*logger.LocalLogger
	c 			*http.Client
	cfg 		*config.OneCloud
}

func NewOneCClient(ll *logger.LocalLogger, cfg *config.OneCloud) * OneCClient  {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	clt := OneCClient{
		log: 	ll,
		c: 		&c,
		cfg: 	cfg,
	}
	return &clt
}


package discord

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"net/http"
	"time"
)

type DiscordClient struct {
	log    		*logger.LocalLogger
	c 			*http.Client
	cfg 		*config.DiscordConfig
}

func NewDiscordClient(ll *logger.LocalLogger, cfg *config.DiscordConfig) *DiscordClient  {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	clt := DiscordClient{
		log: 	ll,
		c: 		&c,
		cfg: 	cfg,
	}
	return &clt
}


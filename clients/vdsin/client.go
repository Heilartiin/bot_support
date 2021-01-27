package vdsin

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"net/http"
	"time"
)

type VDSinClient struct {
	log    		*logger.LocalLogger
	c 			*http.Client
	cfg 		*config.VDSin
}

func NewVDSinClient(ll *logger.LocalLogger, cfg *config.VDSin) *VDSinClient  {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	clt := VDSinClient{
		log: 	ll,
		c: 		&c,
		cfg: 	cfg,
	}
	return &clt
}



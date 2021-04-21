package gs_nike_products

import (
	"github.com/Heilartin/bot_support/logger"
	"net/http"
	"time"
)

type Client struct {
	ApiUrl  string
	c       *http.Client
	log 	*logger.LocalLogger
}

func NewClient(logger *logger.LocalLogger) *Client  {
	c := http.Client{Timeout: 15 * time.Second}
	cl := Client{
		ApiUrl: "https://gs-product.nike.com",
		c:   	&c,
		log: 	logger,
	}
	return &cl
}

package ethereum

import (
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type Client struct {
	log       *logger.LocalLogger
	ethClient *ethClient.Client
	cfg 	  *config.Infura
}

func NewClient(log *logger.LocalLogger, cfg *config.Infura) *Client {
	return &Client{
		log:       log,
		ethClient: InitEthClient(cfg),
		cfg:  	   cfg,
	}
}

func InitEthClient(cfg *config.Infura) *ethClient.Client {
	client, err := ethClient.Dial(cfg.Http)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

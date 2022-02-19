package main

import (
	"github.com/Heilartin/bot_support/clients/ethereum"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/k0kubun/pp"
	"math/big"
)

func main() {
	cfg := config.NewConfig()
	log := logger.NewLogger(cfg.ProductionStart)
	ethClient := ethereum.NewClient(log, cfg.Infura)
	discord := discordgo.Session{}
}


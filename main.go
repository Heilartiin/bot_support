package main

import (
	"github.com/Heilartin/bot_support/app"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println()
	cfg := config.NewConfig()
	ll := logger.NewLogger(cfg.ProductionStart)
	ap, err := app.New(cfg, ll)
	if err != nil {
		ll.Fatal(err)
	}
	ap.Run()
}

//func main() {
//	cfg := config.NewConfig()
//	log := logger.NewLogger(cfg.ProductionStart)
//	ethClient := ethereum.NewClient(log, cfg.Infura)
//
//	response, pending, err := ethClient.TransactionByHash("0x49aa153baf1c773a1ab69f341dfefbb92b7ec2b326cdc95e897c148f407736fa")
//	if err != nil {
//		log.Fatal(err)
//	}
//	_ = pending
//
//	os := opensea.NewClient(log, cfg.OpenSea)
//
//	result, err := os.GetInformationByContract(response.To().String())
//	if err != nil {
//		log.Fatal(err)
//	}
//	pp.Println(result)
//}

package main

//
//func main() {
//	cfg := config.NewConfig()
//	log := logger.NewLogger(cfg.ProductionStart)
//	ethClient := ethereum.NewClient(log, cfg.Infura)
//
//	response, pending, err := ethClient.TransactionByHash("0xa3039b052c95630a9d1e39207769b1d29000f62aa332bbef8a7bc42bd6fe79fd")
//	if err != nil {
//		log.Fatal(err)
//	}
//	_ = pending
//
//	pp.Println(response.Type())
//
//	blockNumber, err := ethClient.GetCurrentBlockNumber()
//	if err != nil {
//		log.Fatal(err)
//	}
//	pp.Println(int(blockNumber))
//	block, err := ethClient.GetBlockByNumber(big.NewInt(int64(13634732)))
//	if err != nil {
//		log.Fatal(err)
//	}
//	//br, err := response.MarshalJSON()
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//fmt.Println(string(br))
//}


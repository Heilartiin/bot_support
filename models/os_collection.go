package models

import "time"

type OpenSeaCollection struct {
	Address             string  `json:"address"`
	EtherscanUrl        string  `json:"contract_url"`
	TxsEtherscan        string  `json:"txs_etherscan"`
	PendingTxsEtherscan string  `json:"pending_txs_etherscan"`
	Name                string  `json:"name"`
	Slug                string  `json:"slug"`
	OSUrl               string  `json:"os_url"`
	ImageUrl            string  `json:"image_url"`
	ServiceFee          float64 `json:"service_fee"`
	CreatorFee          float64 `json:"creator_fee"`

	FloorPrice  float64 `json:"floor_price"`
	TotalVolume float64 `json:"total_volume"`
	TotalSales  float64 `json:"total_sales"`
	NumOwners   int     `json:"num_owners"`

	ExternalLink string `json:"external_link"`
	DiscordUrl   string `json:"discord_url"`
	TelegramUrl  string `json:"telegram_url"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`

	ContractCreated     time.Time `json:"contract_created"`
	OSCollectionCreated time.Time `json:"os_collection_created"`
}

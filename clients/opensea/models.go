package opensea

type (
	ContractResponse struct {
		Collection 					*Collection `json:"collection"`
		Address                     string      `json:"address"`
		AssetContractType           string      `json:"asset_contract_type"`
		CreatedDate                 string      `json:"created_date"`
		Name                        string      `json:"name"`
		NftVersion                  string      `json:"nft_version"`
		OpenseaVersion              interface{} `json:"opensea_version"`
		Owner                       int         `json:"owner"`
		SchemaName                  string      `json:"schema_name"`
		Symbol                      string      `json:"symbol"`
		TotalSupply                 string      `json:"total_supply"`
		Description                 string      `json:"description"`
		ExternalLink                string      `json:"external_link"`
		ImageUrl                    string      `json:"image_url"`
		DefaultToFiat               bool        `json:"default_to_fiat"`
		DevBuyerFeeBasisPoints      float64         `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     float64         `json:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  float64         `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints float64         `json:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
		PayoutAddress               string      `json:"payout_address"`
	}

	Collection struct {
		BannerImageUrl          string      `json:"banner_image_url"`
		ChatUrl                 interface{} `json:"chat_url"`
		CreatedDate             string      `json:"created_date"`
		DefaultToFiat           bool        `json:"default_to_fiat"`
		Description             string      `json:"description"`
		DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
		DiscordUrl              *string      `json:"discord_url"`
		DisplayData             struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalUrl                 *string      `json:"external_url"`
		Featured                    bool        `json:"featured"`
		FeaturedImageUrl            interface{} `json:"featured_image_url"`
		Hidden                      bool        `json:"hidden"`
		SafelistRequestStatus       string      `json:"safelist_request_status"`
		ImageUrl                    string      `json:"image_url"`
		IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
		LargeImageUrl               interface{} `json:"large_image_url"`
		MediumUsername              interface{} `json:"medium_username"`
		Name                        string      `json:"name"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
		PayoutAddress               string      `json:"payout_address"`
		RequireEmail                bool        `json:"require_email"`
		ShortDescription            interface{} `json:"short_description"`
		Slug                        string      `json:"slug"`
		TelegramUrl                 *string `json:"telegram_url"`
		TwitterUsername             *string      `json:"twitter_username"`
		InstagramUsername           *string      `json:"instagram_username"`
		WikiUrl                     *string `json:"wiki_url"`
	}
)

type (
	CollectionResponse struct {
	Collection struct {
		Editors       []string `json:"editors"`
		PrimaryAssetContracts []*PrimaryAssetContracts `json:"primary_asset_contracts"`
		Stats *Stats `json:"stats"`
		BannerImageUrl          string      `json:"banner_image_url"`
		ChatUrl                 interface{} `json:"chat_url"`
		CreatedDate             string      `json:"created_date"`
		DefaultToFiat           bool        `json:"default_to_fiat"`
		Description             string      `json:"description"`
		DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
		DiscordUrl              string      `json:"discord_url"`
		DisplayData             struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalUrl                 string      `json:"external_url"`
		Featured                    bool        `json:"featured"`
		FeaturedImageUrl            interface{} `json:"featured_image_url"`
		Hidden                      bool        `json:"hidden"`
		SafelistRequestStatus       string      `json:"safelist_request_status"`
		ImageUrl                    string      `json:"image_url"`
		IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
		LargeImageUrl               interface{} `json:"large_image_url"`
		MediumUsername              interface{} `json:"medium_username"`
		Name                        string      `json:"name"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
		PayoutAddress               string      `json:"payout_address"`
		RequireEmail                bool        `json:"require_email"`
		ShortDescription            interface{} `json:"short_description"`
		Slug                        string      `json:"slug"`
		TelegramUrl                 interface{} `json:"telegram_url"`
		TwitterUsername             string      `json:"twitter_username"`
		InstagramUsername           string      `json:"instagram_username"`
		WikiUrl                     interface{} `json:"wiki_url"`
	} `json:"collection"`
	}
	PrimaryAssetContracts struct {
		Address                     string      `json:"address"`
		AssetContractType           string      `json:"asset_contract_type"`
		CreatedDate                 string      `json:"created_date"`
		Name                        string      `json:"name"`
		NftVersion                  string      `json:"nft_version"`
		OpenseaVersion              interface{} `json:"opensea_version"`
		Owner                       int         `json:"owner"`
		SchemaName                  string      `json:"schema_name"`
		Symbol                      string      `json:"symbol"`
		TotalSupply                 string      `json:"total_supply"`
		Description                 string      `json:"description"`
		ExternalLink                string      `json:"external_link"`
		ImageUrl                    string      `json:"image_url"`
		DefaultToFiat               bool        `json:"default_to_fiat"`
		DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
		PayoutAddress               string      `json:"payout_address"`
	}
	Stats struct{
		OneDayVolume          float64 `json:"one_day_volume"`
		OneDayChange          float64 `json:"one_day_change"`
		OneDaySales           float64 `json:"one_day_sales"`
		OneDayAveragePrice    float64 `json:"one_day_average_price"`
		SevenDayVolume        float64 `json:"seven_day_volume"`
		SevenDayChange        float64 `json:"seven_day_change"`
		SevenDaySales         float64 `json:"seven_day_sales"`
		SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
		ThirtyDayVolume       float64 `json:"thirty_day_volume"`
		ThirtyDayChange       float64 `json:"thirty_day_change"`
		ThirtyDaySales        float64 `json:"thirty_day_sales"`
		ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
		TotalVolume           float64 `json:"total_volume"`
		TotalSales            float64 `json:"total_sales"`
		TotalSupply           float64 `json:"total_supply"`
		Count                 float64 `json:"count"`
		NumOwners             int     `json:"num_owners"`
		AveragePrice          float64 `json:"average_price"`
		NumReports            int     `json:"num_reports"`
		MarketCap             float64 `json:"market_cap"`
		FloorPrice            float64 `json:"floor_price"`
	}
)

type StatsResponse struct {
	Stats *Stats `json:"stats"`
}

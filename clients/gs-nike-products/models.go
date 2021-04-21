package gs_nike_products

import "time"

type ProductResponse struct {
	ID               	string 					`json:"id"`
	MerchGroup       	string 					`json:"merchGroup"`
	PID              	string 					`json:"pId"`
	Name             	string 					`json:"name"`
	ProductMerchSkus 	[]*ProductSku 			`json:"productMerchSkus"`
	CountryDetails 		[]*ProductCountryDetail `json:"countryDetails"`
	LastRefreshTime 	time.Time     			`json:"lastRefreshTime"`
	StyleColor      	string        			`json:"styleColor"`
}

type ProductSku struct {
	ID                 string        `json:"id"`
	NikeSize           string        `json:"nikeSize"`
	StockKeepingUnitID string        `json:"stockKeepingUnitId"`
	Gtin               string        `json:"gtin"`
}

type ProductCountryDetail struct {
	CountryCode        		string        		`json:"countryCode"`
	LangLocale         		string        		`json:"langLocale"`
	CanSell            		bool          		`json:"canSell"`
	DescriptionHTML    		string        		`json:"descriptionHtml"`
	FullTitle          		string        		`json:"fullTitle"`
	Title              		string        		`json:"title"`
	Subtitle           		string        		`json:"subtitle"`
	DescriptionHeading 		string       		`json:"descriptionHeading"`
	ColorDescription   		string        		`json:"colorDescription"`
	Charges            		[]*ProductCharge 	`json:"charges"`
	ProductImageURL      	string        		`json:"productImageUrl"`
	IsNationalIDRequired 	bool          		`json:"isNationalIDRequired"`
}

type ProductCharge struct {
	Name  string `json:"name"`
	Value *ProductChargeValue `json:"value"`
}

type ProductChargeValue struct {
	CurrencyCode string `json:"currencyCode"`
	Value        float64    `json:"value"`
}
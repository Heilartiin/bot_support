package mrporter

import "time"

type MrpResp struct {
	RecordSetTotal        	int    				`json:"recordSetTotal"`
	ResourceID            	string 				`json:"resourceId"`
	RecordSetCount        	int    				`json:"recordSetCount"`
	RecordSetComplete     	string 				`json:"recordSetComplete"`
	RecordSetStartNumber  	int    				`json:"recordSetStartNumber"`
	RecordSetTotalMatches 	int    				`json:"recordSetTotalMatches"`
	Products             	[]*MrpProduct  		`json:"products"`
	ResourceName 			string  			`json:"resourceName"`
	Version      			float64 			`json:"version"`
}

type MrpProduct struct {
	Dynamic     bool `json:"dynamic"`
	Visible     bool `json:"visible"`
	DesignerSeo struct {
		SeoURLKeyword string `json:"seoURLKeyword"`
	} `json:"designerSeo"`
	Displayable              bool     `json:"displayable"`
	DesignerNameEN           string   `json:"designerNameEN"`
	Type                     string   `json:"type"`
	ExternalReccomendationID []string `json:"externalReccomendationId"`
	Name                     string   `json:"name"`
	DesignerIdentifier       string   `json:"designerIdentifier"`
	ForceLogIn               bool     `json:"forceLogIn"`
	MfPartNumber             string   `json:"mfPartNumber"`
	PartNumber               string   `json:"partNumber"`
	ProductColours           []struct {
		Visible                  bool     `json:"visible"`
		EditorialDescription     string   `json:"editorialDescription"`
		DetailsAndCare           string   `json:"detailsAndCare"`
		Displayable              bool     `json:"displayable"`
		Type                     string   `json:"type"`
		ExternalReccomendationID []string `json:"externalReccomendationId"`
		Swatch                   struct {
			URL string `json:"URL"`
		} `json:"swatch"`
		TechnicalDescription             string    `json:"technicalDescription"`
		Selected                         bool      `json:"selected"`
		WhatsNewStart3074457345616676721 string    `json:"whats_new_start_3074457345616676721"`
		IsDefault                        bool      `json:"isDefault"`
		ForceLogIn                       bool      `json:"forceLogIn"`
		MfPartNumber                     string    `json:"mfPartNumber"`
		PartNumber                       string    `json:"partNumber"`
		ImageViews                       []string  `json:"imageViews"`
		FirstVisibleDate                 time.Time `json:"firstVisibleDate"`
		SKUs                             []struct {
			SkuUniqueID      int64  `json:"skuUniqueID"`
			NotStockedOnline bool   `json:"notStockedOnline,omitempty"`
			Displayable      bool   `json:"displayable"`
			Type             string `json:"type"`
			Banned           bool   `json:"banned"`
			Size             struct {
				CentralSizeLabel string `json:"centralSizeLabel"`
				Schemas          []struct {
					Name     string   `json:"name"`
					Labels   []string `json:"labels"`
					Selected bool     `json:"selected,omitempty"`
				} `json:"schemas"`
				ScaleLabel string `json:"scaleLabel"`
				LabelSize  string `json:"labelSize"`
			} `json:"size"`
			Selected bool `json:"selected"`
			Price    struct {
				SellingPrice struct {
					Amount  int `json:"amount"`
					Divisor int `json:"divisor"`
				} `json:"sellingPrice"`
				RdSellingPrice struct {
					Amount  int `json:"amount"`
					Divisor int `json:"divisor"`
				} `json:"rdSellingPrice"`
				Currency struct {
					Symbol string `json:"symbol"`
					Label  string `json:"label"`
				} `json:"currency"`
			} `json:"price"`
			Composition string `json:"composition"`
			Buyable     bool   `json:"buyable"`
			ForceLogIn  bool   `json:"forceLogIn"`
			Attributes  []struct {
				Values []struct {
					Label      string `json:"label"`
					Identifier string `json:"identifier"`
				} `json:"values"`
				Usage      string `json:"usage"`
				Label      string `json:"label"`
				Identifier string `json:"identifier"`
			} `json:"attributes"`
			PartNumber    string `json:"partNumber"`
			SoldOutOnline bool   `json:"soldOutOnline,omitempty"`
			Badges        []struct {
				Label string `json:"label"`
				Type  string `json:"type"`
				Key   string `json:"key"`
			} `json:"badges,omitempty"`
		} `json:"sKUs"`
		Label         string `json:"label"`
		Banned        bool   `json:"banned"`
		ProductID     int64  `json:"productId"`
		SoldOutOnline bool   `json:"soldOutOnline"`
		Badges        []struct {
			Label string `json:"label"`
			Type  string `json:"type"`
			Key   string `json:"key"`
		} `json:"badges"`
		Price struct {
			SellingPrice struct {
				Amount  int `json:"amount"`
				Divisor int `json:"divisor"`
			} `json:"sellingPrice"`
			RdSellingPrice struct {
				Amount  int `json:"amount"`
				Divisor int `json:"divisor"`
			} `json:"rdSellingPrice"`
			Currency struct {
				Symbol string `json:"symbol"`
				Label  string `json:"label"`
			} `json:"currency"`
		} `json:"price"`
		ImageTemplate    string `json:"imageTemplate"`
		ShortDescription string `json:"shortDescription"`
		MadeIn           string `json:"madeIn"`
		Buyable          bool   `json:"buyable"`
		Seo              struct {
			SeoURLKeyword string `json:"seoURLKeyword"`
		} `json:"seo"`
		Attributes []struct {
			Values []struct {
				Label      string `json:"label"`
				Identifier string `json:"identifier"`
			} `json:"values"`
			Usage      string `json:"usage"`
			Label      string `json:"label"`
			Identifier string `json:"identifier"`
		} `json:"attributes"`
		Identifier string `json:"identifier"`
	} `json:"productColours"`
	SizeAndFit        string    `json:"sizeAndFit"`
	CentralSizeScheme string    `json:"centralSizeScheme"`
	FirstVisibleDate  time.Time `json:"firstVisibleDate"`
	MasterCategory    struct {
		Child struct {
			Child struct {
				CategoryID string `json:"categoryId"`
				Label      string `json:"label"`
				Identifier string `json:"identifier"`
			} `json:"child"`
			CategoryID string `json:"categoryId"`
			Label      string `json:"label"`
			Identifier string `json:"identifier"`
		} `json:"child"`
		CategoryID string `json:"categoryId"`
		Label      string `json:"label"`
		Identifier string `json:"identifier"`
	} `json:"masterCategory"`
	ProductID     string `json:"productId"`
	SoldOutOnline bool   `json:"soldOutOnline"`
	Badges        []struct {
		Label string `json:"label"`
		Type  string `json:"type"`
		Key   string `json:"key"`
	} `json:"badges"`
	Price struct {
		SellingPrice struct {
			Amount  int `json:"amount"`
			Divisor int `json:"divisor"`
		} `json:"sellingPrice"`
		RdSellingPrice struct {
			Amount  int `json:"amount"`
			Divisor int `json:"divisor"`
		} `json:"rdSellingPrice"`
		Currency struct {
			Symbol string `json:"symbol"`
			Label  string `json:"label"`
		} `json:"currency"`
	} `json:"price"`
	Thumbnail string `json:"thumbnail"`
	Tracking  struct {
		DesignerName string `json:"designerName"`
		Name         string `json:"name"`
	} `json:"tracking"`
	DesignerName    string   `json:"designerName"`
	Buyable         bool     `json:"buyable"`
	Recommendations []string `json:"recommendations"`
	Seo             struct {
		Title           string `json:"title"`
		MetaDescription string `json:"metaDescription"`
		MetaKeyword     string `json:"metaKeyword"`
		SeoURLKeyword   string `json:"seoURLKeyword"`
	} `json:"seo"`
	Attributes []struct {
		Values []struct {
			Label      string `json:"label"`
			Identifier string `json:"identifier"`
		} `json:"values"`
		Usage      string `json:"usage"`
		Label      string `json:"label"`
		Identifier string `json:"identifier"`
	} `json:"attributes"`
}
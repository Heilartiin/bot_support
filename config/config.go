package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser          string
	DBPass          string
	DBName          string
	DBSchema        string
	DBHost          string
	DBPort          string
	ProductionStart string
	ProxyMarket     *ProxyMarket
	VDSin           *VDSin
	OneCloud        *OneCloud
	Discord         *Discord
	MRPConfig       *MrpPorter
	DiscordConfig   *DiscordConfig
	NFTDiscord      *NFTDiscord
	OpenSea         *OpenSea
	Infura          *Infura
}


type Discord struct {
	ClientID 		string
	BotToken		string
	PublicKey 		string
	ClientSecret 	string
	Prefix			string
	FooterIcon 		string
}

type OpenSea struct {
	ApiKey 			string
	ApiUrl 			string
	UserAgent 		string
}

type MrpPorter struct {
	ApiUrl 		string
	ClientID 	string
}

type ProxyMarket struct {
	ApiKey string
}

type DiscordConfig struct {
	Url    		string
	Color  		string
	DefaultIMG 	string
	FooterIMG   string
	UserName    string
	QTUrl		string
	Avatar 		string
}

type VDSin struct {
	ApiUrl	string
}

type OneCloud struct {
	ApiUrl 	string
}

type NFTDiscord struct {
	ApiUrl 		string
	UserAgent 	string
}

type Infura struct {
	ProjectID     string
	ProjectSecret string
	Http          string
	Wss           string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[WARNING] no .env file, reading config from OS ENV variables, error: ", err)
	}
	cfg := Config{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBSchema: os.Getenv("DB_SCHEMA"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		ProductionStart: os.Getenv("PRODUCTION_START"),
		Discord: &Discord{
			ClientID: os.Getenv("DISCORD_CLIENT_ID"),
			BotToken: os.Getenv("DISCORD_BOT_TOKEN"),
			PublicKey: os.Getenv("DISCORD_PUBLIC_KEY"),
			ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			Prefix: os.Getenv("DISCORD_PREFIX"),
			FooterIcon: os.Getenv("DISCORD_FOOTER_ICON"),
		},
		MRPConfig: &MrpPorter{
			ApiUrl: os.Getenv("MRPORTER_API_URL"),
			ClientID: os.Getenv("CLIENT_ID"),
		},
		DiscordConfig: &DiscordConfig{
			Url:        os.Getenv("DISCORD_URL"),
			Color:      os.Getenv("DISCORD_COLOR"),
			DefaultIMG: os.Getenv("DISCORD_DEFAULT_IMAGE"),
			FooterIMG:  os.Getenv("DISCORD_FOOTER_ICON"),
			UserName:   os.Getenv("DISCORD_USERNAME"),
			QTUrl:      os.Getenv("QT_API_URL"),
			Avatar:     os.Getenv("DISCORD_AVATAR"),
		},
		VDSin: &VDSin{ApiUrl: os.Getenv("VDSIN_API_URL")},
		OneCloud: &OneCloud{ApiUrl: os.Getenv("ONE_CLOUD_API_URL")},
		ProxyMarket: &ProxyMarket{ApiKey: os.Getenv("PROXY_MARKET_API_KEY")},
		NFTDiscord: &NFTDiscord{
			ApiUrl:    "https://discord.com",
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
		},
		OpenSea: &OpenSea{
			ApiKey: "38fd11fa977640a9be1c01cdc6f43785",
			ApiUrl: "https://api.opensea.io",
			UserAgent: "MetaMask/795 CFNetwork/1312 Darwin/21.0.0",
		},
		Infura: &Infura{
			ProjectID:     "40a8613c773944a9b0f0fa09560038ae",
			ProjectSecret: "2e026e62d9e540d58aafd0f7bafbdb54",
			Http:          "https://mainnet.infura.io/v3/40a8613c773944a9b0f0fa09560038ae",
			Wss:           "wss://mainnet.infura.io/ws/v3/40a8613c773944a9b0f0fa09560038ae",
		},
	}
	return &cfg
}

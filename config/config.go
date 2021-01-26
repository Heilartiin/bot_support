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
	Discord			*Discord
	MRPConfig		*MrpPorter
	DiscordConfig   *DiscordConfig
}


type Discord struct {
	ClientID 		string
	BotToken		string
	PublicKey 		string
	ClientSecret 	string
	Prefix			string
	FooterIcon 		string
}

type MrpPorter struct {
	ApiUrl 		string
	ClientID 	string
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
	}
	return &cfg
}
package controllers

import (
	_cloud "github.com/Heilartin/bot_support/clients/1cloud"
	"github.com/Heilartin/bot_support/clients/discord"
	gsp "github.com/Heilartin/bot_support/clients/gs-nike-products"
	"github.com/Heilartin/bot_support/clients/mrporter"
	"github.com/Heilartin/bot_support/clients/opensea"
	"github.com/Heilartin/bot_support/clients/proxies"
	"github.com/Heilartin/bot_support/clients/vdsin"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/storage"
	proxy_market "github.com/Heilartin/proxy-market"
	"github.com/bwmarrin/discordgo"
	"time"
)


type Controllers struct {
	StartTime     time.Time
	Logger        *logger.LocalLogger
	Config        *config.Config
	Repository    *storage.Storage
	Session 	  *discordgo.Session
	FileClient 	  *proxies.FClient
	Dis 		  *discord.DiscordClient
	MrPorter 	  *mrporter.MrpClient
	VDSin 		  *vdsin.VDSinClient
	OneCloud 	  *_cloud.OneCClient
	OpenSea 	  *opensea.Client
	GsProductNike *gsp.Client
	ProxiesMarket *proxy_market.ProxyMarketClient
}

func NewControllers(rep *storage.Storage, log *logger.LocalLogger,
	cfg *config.Config, s *discordgo.Session) *Controllers {

	a := &Controllers{
		StartTime:     time.Now(),
		Config:        cfg,
		Logger:        log,
		Repository:    rep,
		Session: 	   s,
		VDSin:         vdsin.NewVDSinClient(log, cfg.VDSin),
		FileClient:    proxies.NewFileClient(log, rep),
		OpenSea:       opensea.NewClient(log, cfg.OpenSea),
		MrPorter:  	   mrporter.NewMrpClient(log, cfg.MRPConfig, rep),
		Dis:  		   discord.NewDiscordClient(log, cfg.DiscordConfig),
		OneCloud:      _cloud.NewOneCClient(log, cfg.OneCloud),
		ProxiesMarket: proxy_market.NewProxyMarketClient(cfg.ProxyMarket.ApiKey),
	}
	return a
}

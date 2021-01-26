package app

import (
	"github.com/Heilartin/bot_support/api"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/storage"
	"github.com/bwmarrin/discordgo"
	"time"
)

// App struct is base struct with all essential information about application
type App struct {
	StartTime time.Time
	Logger    *logger.LocalLogger
	Config    *config.Config
}

// New inits new App instance
func New(cfg *config.Config, l *logger.LocalLogger) (*App, error) {
	return &App{
		Logger:    l,
		Config:    cfg,
		StartTime: time.Now(),
	}, nil
}

func (a *App) Run() {
	st, err := storage.NewStorage(a.Config, a.Logger)
	if err != nil {
		a.Logger.Fatalln("can't create new storage: ", err)
	}
	session, _ := discordgo.New()
	session.Token = "Bot " + a.Config.Discord.BotToken
	api.NewAPI(st, a.Logger, a.Config, session).ServeAPI()
}

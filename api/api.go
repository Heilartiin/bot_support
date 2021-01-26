package api

import (
	"github.com/Heilartin/bot_support/api/controllers"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/Heilartin/bot_support/storage"
	"github.com/bwmarrin/discordgo"
	"github.com/bwmarrin/disgord/x/mux"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type API struct {
	*controllers.Controllers
	StartTime  time.Time
	Logger     *logger.LocalLogger
	Config     *config.Config
	Repository *storage.Storage
	Router     *mux.Mux
}

func NewAPI(rep *storage.Storage, log *logger.LocalLogger, cfg *config.Config, sess *discordgo.Session) *API {
	a := &API{
		controllers.NewControllers(rep, log, cfg, sess),
		time.Now(),
		log,
		cfg,
		rep,
		mux.New(),
	}
	return a
}

func (api *API) ServeAPI() {
	// Open a websocket connection to Discord
	err := api.Session.Open()
	if err != nil {
		api.Logger.Fatalln("error opening connection to Discord, %s\n", err)
	}
	api.Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	api.Session.AddHandler(api.MessageHandler)
	api.Router.Route("help", "Display this message.", api.Router.Help)

	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	api.Session.Close()
}


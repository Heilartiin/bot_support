package api

import (
	"context"
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
)

//667588661470035989
func (api *API) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate)  {

	// Ignore all message prefix != "!"
	if !strings.HasPrefix(m.Content, api.Config.Discord.Prefix) {
		return
	}
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Split command without prefix
	content := strings.Replace(m.Content, api.Config.Discord.Prefix, "", 1)
	command := strings.Fields(content)

	if len(command) == 0 {
		return
	}

	if command[0] == "auth" {
		err := api.Controllers.CreateUser(m)
		if err != nil {
			return
		}
		return
	}

	user, err := api.Repository.DB.GetUserByMemberID(m.Author.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			api.BadAction("Sorry, we could not authenticate you. Try again", m)
			return
		}
		api.Logger.Error(errors.WithStack(err))
		return
	}

	ctx := context.WithValue(context.Background(), "userData", user)

	switch command[0] {
	case "atc":
		api.Controllers.GetAtc(ctx, m)
	case "atw":
		api.Controllers.GetAtw(ctx, m)
	case "qt":
		api.Controllers.GetQTs(m)
		default:
			return
		}
}



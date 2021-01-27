package controllers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"time"
)

func (c *Controllers) BadAction(description string, m *discordgo.MessageCreate)  {
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	user := discordgo.MessageEmbedField{
		Name: "User",
		Value: fmt.Sprintf("%s#%s (<@%s>)", m.Author.Username, m.Author.Discriminator, m.Author.ID),
		Inline: true,
	}
	e := discordgo.MessageEmbed{
		Title: "Error",
		Description: fmt.Sprintf("> " + description),
		Fields: []*discordgo.MessageEmbedField{&user},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Manager | %v", time.Now().Format("Mon Jan _2 15:04:05 2006")),
			IconURL: c.Config.Discord.FooterIcon,
		},
	}
	_, err = c.Session.ChannelMessageSendEmbed(channelPrivate.ID, &e)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}


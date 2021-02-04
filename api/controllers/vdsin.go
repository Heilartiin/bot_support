package controllers

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"

  	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (c *Controllers) VDSinGetBalance(m *discordgo.MessageCreate)  {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	apiToken := info[1]

	user := discordgo.MessageEmbedField{
		Name: "User",
		Value: fmt.Sprintf("%s#%s (<@%s>)", m.Author.Username, m.Author.Discriminator, m.Author.ID),
		Inline: false,
	}
	balance, err := c.VDSin.GetBalance(apiToken)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	e := discordgo.MessageEmbed{
		Title: "VDSin Balance",
		Description: fmt.Sprintf("Your balance: %s", balance.Data.Real),
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


func (c *Controllers) VDSinGetAllServers(m *discordgo.MessageCreate)  {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	apiToken := info[1]
	resp, err := c.VDSin.GetAllCredential(apiToken)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	csvContent, err := gocsv.MarshalString(&resp) // Get all clients as CSV string
	//err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	br := []byte(csvContent)
	nameFile := fmt.Sprintf("vd-servers-%s.csv", time.Now().Format("15:04:05"))
	_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, bytes.NewBuffer(br))
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}


func (c *Controllers) VDSinDeleteAllServers(m *discordgo.MessageCreate)  {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	apiToken := info[1]
	err := c.VDSin.DeleteAllServers(apiToken)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	c.SuccessAction("All server of VDSin deleted.", m)
	return
}



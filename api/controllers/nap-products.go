package controllers

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"time"
)

func (c *Controllers) GetNapProductsInvisible(m *discordgo.MessageCreate)  {
	resp, err := c.Repository.DB.GetNapProductInvisible()
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	csvContent, err := gocsv.MarshalString(&resp)
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
	nameFile := fmt.Sprintf("invisible-products-%s.csv", time.Now().Format("15:04:05"))
	_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, bytes.NewBuffer(br))
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	return
}

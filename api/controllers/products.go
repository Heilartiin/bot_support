package controllers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
)
var stores = []string{"mrp_RU", "mrp_GB", "mrp_US"}

func (c *Controllers) GetQTs(m *discordgo.MessageCreate) {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	pid := info[1]
	for _, store := range stores {
		resp, err := c.Repository.DB.GetProductByTaskID(pid, store)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			c.BadAction("Bad request", m)
			return
		}
		err = c.Dis.SendBotInfoNotification(resp)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			c.BadAction("Bad request", m)
			return
		}
	}
	return
}

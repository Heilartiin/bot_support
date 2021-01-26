package controllers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Heilartin/bot_support/clients/mrporter"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (c *Controllers) GetAtc(ctx context.Context, m *discordgo.MessageCreate) {
	userInfo, err := c.UserFromContext(ctx)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	pid := info[1]
	var product *mrporter.MrpProduct
	for _, storeID := range models.Stores {
		resp, err := c.MrPorter.GetProductInfo(storeID, pid)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			continue
		}
		product = resp
		break
	}
	e := c.CreateATCFile(product, userInfo)
	b := []byte(e)
	_, err = c.Session.ChannelFileSend(userInfo.PrivateChannel, fmt.Sprintf("atc-%s-%s-%v.html", pid, userInfo.MemberID, time.Now().Unix()), bytes.NewBuffer(b))
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}

func (c *Controllers) GetAtw(ctx context.Context, m *discordgo.MessageCreate) {
	userInfo, err := c.UserFromContext(ctx)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	info := strings.Split(m.Content, " ")
	if len(info) < 1 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	pid := info[1]
	var product *mrporter.MrpProduct
	for _, storeID := range models.Stores {
		resp, err := c.MrPorter.GetProductInfo(storeID, pid)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			continue
		}
		product = resp
		break
	}
	resp, err := c.CreateATWBody(product)
	_, err = c.Session.ChannelFileSend(userInfo.PrivateChannel, fmt.Sprintf("atw.json"), resp)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}
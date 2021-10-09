package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"time"
)

func (c *Controllers) GetProxyMarketProxiesJSONFile(m *discordgo.MessageCreate)  {
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	file, err := c.ProxiesMarket.GetProxyListAllByNewest()
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	f, err := json.Marshal(file.List.Data)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, f,"", "\t")
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	nameFile := fmt.Sprintf("proxies-json-%s.json", time.Now().Format("15:04:05"))
	_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, &prettyJSON)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	return
}

func (c *Controllers) GetProxyMarketStringJSONFile(m *discordgo.MessageCreate)  {
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	file, err := c.ProxiesMarket.GetProxyListAllByNewest()
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	var prxs []string
	for _, v := range file.List.Data {
		if v.ExpiredAt.Time.Before(time.Now()) {
			continue
		}
		prxs = append(prxs, fmt.Sprintf("%s:%s:%s:%s",
			v.IP, v.HTTPPort, v.Login, v.Password))
	}
	f, err := json.Marshal(prxs)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, f,"", "\t")
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	nameFile := fmt.Sprintf("proxies-string-%s.json", time.Now().Format("15:04:05"))
	_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, &prettyJSON)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	return
}
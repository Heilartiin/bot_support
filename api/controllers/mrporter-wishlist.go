package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Heilartin/bot_support/clients/mrporter"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
)

func (c *Controllers) CreateATWBody(products *mrporter.MrpProduct) (*bytes.Buffer, error) {
	var partNumbers []*PartNumber
	for _, v := range products.ProductColours[0].SKUs {
		partNumbers = append(partNumbers, &PartNumber{PartNumber: v.PartNumber})
	}
	resp := ItemATW{Item: partNumbers}
	br, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, br,"", "\t")
	if err != nil {
		return nil, err
	}
	return &prettyJSON, nil
}

type ItemATW struct {
	Item []*PartNumber	`json:"item"`
}

type PartNumber struct {
	PartNumber string `json:"partNumber"`
}


func (c *Controllers) GetAtw(m *discordgo.MessageCreate) {
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
	privateChannel, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	resp, err := c.CreateATWBody(product)
	_, err = c.Session.ChannelFileSend(privateChannel.ID, fmt.Sprintf("atw.json"), resp)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}
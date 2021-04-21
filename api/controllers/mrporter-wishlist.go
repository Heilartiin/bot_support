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
	SizeName   string `json:"sizeName"`
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
			c.Logger.Error(err)
			c.BadAction(err.Error(), m)
			continue
		}
		product = resp
		break
	}
	privateChannel, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	resp, err := c.CreateATWBody(product)
	_, err = c.Session.ChannelFileSend(privateChannel.ID, fmt.Sprintf("atw.json"), resp)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
}


var ValidSizes = map[string]string{
	"4": "4",
	"4.5": "4.5",
	"5": "5",
	"5.5": "5.5",
	"6": "6",
	"6.5": "6.5",
	"7": "7",
	"7.5": "7.5",
	"8": "8",
	"8.5": "8.5",
	"9": "9",
	"9.5": "9.5",
	"10": "10",
	"10.5": "10.5",
	"11": "11",
	"11.5": "11.5",
	"12": "12",
	"12.5": "12.5",
	"13": "13",
	"13.5": "13.5",
}

func (c *Controllers) GetATWFromScraper(m *discordgo.MessageCreate)  {
	info := strings.Split(m.Content, " ")
	if len(info) < 1 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	pid := info[1]
	products, err := c.Repository.DB.GetScrapersProductByPid(pid)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	privateChannel, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}

	var partNumbers []*PartNumber
	for _, v := range products {
		if ValidSizes[v.Size] != "" {
			partNumbers = append(partNumbers,
				&PartNumber{
					PartNumber: v.PartNumber,
					SizeName:   v.Size},
			)
		}
	}
	resp := ItemATW{Item: partNumbers}
	br, err := json.Marshal(resp)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, br,"", "\t")
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	_, err = c.Session.ChannelFileSend(privateChannel.ID, fmt.Sprintf("%s_atw.json", pid), &prettyJSON)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}

//var ValidSizes = []string{
//	"4", "4.5", "5", "6", "6.5",
//	"7", "7.5", "8", "8.5", "9",
//	"9.5", "10", "10.5", "11",
//	"11.5", "12", "12.5", "13",
//}
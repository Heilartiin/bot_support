package controllers

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Heilartin/bot_support/clients/mrporter"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strconv"
	"strings"
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
			Text: fmt.Sprintf("BS | %v", time.Now().Format("Mon Jan _2 15:04:05 2006")),
			IconURL: c.Config.Discord.FooterIcon,
		},
	}
	_, err = c.Session.ChannelMessageSendEmbed(channelPrivate.ID, &e)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}

func (c *Controllers) CreateUserEmbed(title string, u *models.User) *discordgo.MessageEmbed {
	userID := discordgo.MessageEmbedField{
		Name:   "UserID",
		Value: 	strconv.Itoa(u.ID),
		Inline: false,
	}
	wishlist := discordgo.MessageEmbedField{
		Name:   "WishList ID",
		Value: 	u.WishList,
		Inline: false,
	}
	token := discordgo.MessageEmbedField{
		Name:   "Token",
		Value: 	fmt.Sprintf("```%s...........%s```", u.Token[0:15], u.Token[len(u.Token)-15:]),
		Inline: true,
	}
	fields := []*discordgo.MessageEmbedField{&userID, &wishlist, &token}
	e := discordgo.MessageEmbed{
		Title: title,
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Atc Support | %v", time.Now().Format("Mon Jan _2 15:04:05")),
			IconURL: c.Config.Discord.FooterIcon,
		},
	}
	return &e
}

func (c *Controllers) CreateATCFile(products *mrporter.MrpProduct, u *models.User) string {
	var file []string
	file = append(file, "<html>\n<body>")
	for _, storeID := range models.Stores {
		for _, color := range products.ProductColours {
			for _, size := range color.SKUs {
				sEnc := b64.StdEncoding.EncodeToString([]byte(u.Token))
				text := fmt.Sprintf("%s | %s | %s", storeID, size.PartNumber, size.Size.CentralSizeLabel)
				link := fmt.Sprintf("http://heilartin.pw/api/atc/mrp_atc.html?token=%s&pid=%s&store=%s&quantity=1", sEnc, size.PartNumber, storeID)
				htmlStr := fmt.Sprintf("<p><a href=\"%s\">%s</></p>", link, text)
				file = append(file, htmlStr)
			}
		}
	}
	file = append(file, "<html>\n<body>")
	return strings.Join(file[:],"\n")
}

func (c *Controllers) CreateATWFile(products *mrporter.MrpProduct, u *models.User, wishlist string) string {
	var file []string
	file = append(file, "<html>\n<body>")
	for _, storeID := range models.Stores {
		for _, color := range products.ProductColours {
			for _, size := range color.SKUs {
				sEnc := b64.StdEncoding.EncodeToString([]byte(u.Token))
				text := fmt.Sprintf("%s | %s | %s", storeID, size.PartNumber, size.Size.CentralSizeLabel)
				link := fmt.Sprintf("http://heilartin.pw/api/atc/mrp_atw.html?token=%s&pid=%s&store=%s&wish=%s", sEnc, size.PartNumber, storeID, wishlist)
				htmlStr := fmt.Sprintf("<p><a href=\"%s\">%s</></p>", link, text)
				file = append(file, htmlStr)
			}
		}
	}
	file = append(file, "<html>\n<body>")
	return strings.Join(file[:],"\n")
}

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
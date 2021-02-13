package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (d *DiscordClient) doReq(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		d.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	req.Header.Add("Content-type", "application/json")
	resp, err := d.c.Do(req)
	if err != nil {
		d.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return resp, nil
}


func (d *DiscordClient) SendBotInfoNotification(p *models.Product) error {
	e := d.CreateBotEmbed(p)
	w := discordgo.WebhookParams{
		AvatarURL: d.cfg.Avatar,
		Username: d.cfg.UserName,
		Embeds: []*discordgo.MessageEmbed{e},
	}
	b, err := json.Marshal(w)
	if err != nil {
		d.log.Error(errors.WithStack(err))
		return nil
	}
	resp, err := d.doReq(d.cfg.Url, b)
	if err != nil {
		d.log.Error(errors.WithStack(err))
		return nil
	}
	if resp.StatusCode != 204 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			d.log.Error(errors.WithStack(err))
			return nil
		}
		d.log.Error(errors.WithStack(errors.New(fmt.Sprintf("Bad status code: %s ; %s", resp.Status, string(b)))))
		return nil
	}
	return nil
}



func (d *DiscordClient) CreateBotEmbed(p *models.Product) *discordgo.MessageEmbed  {
	pathUrl := ""
	region := ""
	if p.StoreID == "mrp_US"  {
		region = "US"
		pathUrl = "en-us"
	}
	if p.StoreID == "mrp_RU" {
		region = "RU"
		pathUrl = "en-ru"
	}
	if p.StoreID == "mrp_GB" {
		region = "GB"
		pathUrl = "en-gb"
	}

	taskID := discordgo.MessageEmbedField{
		Name:   "Task ID",
		Value: 	fmt.Sprintf("%d", p.Task.ID),
		Inline:  true,
	}
	accessLink := discordgo.MessageEmbedField{
		Name: 	 "Access",
		Value:	 fmt.Sprintf("[Click](https://www.mrporter.com/%s/wishlist/%s/%s)", pathUrl, p.WishListID, p.AccessKey),
		Inline:  true,
	}
	delay := discordgo.MessageEmbedField{
		Name:    "Delay",
		Value: 	 fmt.Sprintf("%d", p.Task.TimeSleep),
		Inline:  true,
	}

	itemCode := discordgo.MessageEmbedField{
		Name:   "Item Code",
		Value: 	 p.Pid,
		Inline: true,
	}
	storeID := discordgo.MessageEmbedField{
		Name: 	"Store ID",
		Value:	p.StoreID,
		Inline: true,
	}
	price := discordgo.MessageEmbedField{
		Name: "Price",
		Value: fmt.Sprintf("%d %s", p.Price / 100, p.Symbol),
		Inline: true,
	}
	fields := []*discordgo.MessageEmbedField{&taskID, &accessLink, &delay, &itemCode, &storeID, &price}

	var links 	[]string
	var quicks 	[]string
	var slow 	[]string

	sizes := p.Sizes
	if len(p.Sizes) > 18 {
		sizes = p.Sizes[0:18]
	}

	for _, v := range sizes {
		text := fmt.Sprintf("> %s %s", v.SizeChart, v.SizeName)
		sizeText := fmt.Sprintf("%s ``%s:1``", text, v.PartNumber)

		fastText := fmt.Sprintf("[quick pay](%s/quicktask/?store=%s&sku=%s:1&mode=4) | [preload](%s/quicktask/?store=%s&sku=%s:1&mode=2)",
			d.cfg.QTUrl, p.StoreID, v.PartNumber, d.cfg.QTUrl, p.StoreID, v.PartNumber)

		slowText := fmt.Sprintf("[default](%s/quicktask/?store=%s&sku=%s:1&mode=1)", d.cfg.QTUrl, p.StoreID, v.PartNumber)

		slow = append(slow, slowText)
		quicks = append(quicks, fastText)
		links = append(links, sizeText)
	}
	size := 5
	var j int
	for i := 0; i < len(links); i += size{
		j += size
		if j > len(links) {
			j = len(links)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		smallSlice := links[i:j]
		slowSlice := slow[i:j]
		quicksSlice := quicks[i:j]

		newSizesField := discordgo.MessageEmbedField{
			Name: "Sizes",
			Value: strings.Join(smallSlice[:],"\n"),
			Inline: true,
		}
		newSkusField := discordgo.MessageEmbedField{
			Name: "Fast-Mode",
			Value: strings.Join(quicksSlice[:],"\n"),
			Inline: true,
		}
		newQuickField := discordgo.MessageEmbedField{
			Name: "Slow-Mode",
			Value: strings.Join(slowSlice[:],"\n"),
			Inline: true,
		}
		fields = append(fields, &newSizesField, &newSkusField, &newQuickField)
	}
	color, _ := strconv.Atoi(d.cfg.Color)
	e := discordgo.MessageEmbed{
		Title: fmt.Sprintf("[%s] %s", region, p.Name),
		URL: fmt.Sprintf("https://www.mrporter.com/%s/product/%s", pathUrl, p.Pid),
		Fields: fields,
		Color: color,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s | %v", "MrPorter QT", time.Now().Format("2006-01-02 15:04:05.000000")),
			IconURL: d.cfg.FooterIMG,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: p.Image,
		},
	}
	return &e
}
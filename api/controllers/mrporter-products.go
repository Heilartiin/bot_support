package controllers

import (
	"database/sql"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
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
		task, err := c.Repository.DB.GetTaskByID(resp.TaskID)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			c.BadAction("Bad request", m)
			return
		}
		resp.Task = task

		embedProduct := c.CreateProductEmbed(resp)
		_, err = c.Session.ChannelMessageSendEmbed(m.ChannelID, embedProduct)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			return
		}
	}
	return
}

func (c *Controllers) GetQTSToPrivateChannel(m *discordgo.MessageCreate) {
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
			if err == sql.ErrNoRows {
				c.Logger.Error(errors.WithStack(err))
				c.BadAction("This is pid not added to database", m)
				return
			}
			c.Logger.Error(errors.WithStack(err))
			c.BadAction("Bad request", m)
			return
		}
		task, err := c.Repository.DB.GetTaskByID(resp.TaskID)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			c.BadAction("Bad request", m)
			return
		}
		resp.Task = task
		embedProduct := c.CreateProductEmbed(resp)

		privateChannel, err := c.Session.UserChannelCreate(m.Author.ID)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			return
		}
		_, err = c.Session.ChannelMessageSendEmbed(privateChannel.ID, embedProduct)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			return
		}
	}
	return
}


func (c *Controllers) CreateProductEmbed(p *models.Product) *discordgo.MessageEmbed  {
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
			c.Config.DiscordConfig.QTUrl, p.StoreID, v.PartNumber, c.Config.DiscordConfig.QTUrl, p.StoreID, v.PartNumber)

		slowText := fmt.Sprintf("[default](%s/quicktask/?store=%s&sku=%s:1&mode=1)",  c.Config.DiscordConfig.QTUrl, p.StoreID, v.PartNumber)

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

	color, _ := strconv.Atoi(c.Config.DiscordConfig.Color)
	e := discordgo.MessageEmbed{
		Title: fmt.Sprintf("[%s] %s", region, p.Name),
		URL: fmt.Sprintf("https://www.mrporter.com/%s/product/%s", pathUrl, p.Pid),
		Fields: fields,
		Color: color,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s | %v", "MrPorter QT", time.Now().Format("2006-01-02 15:04:05.000000")),
			IconURL:c.Config.DiscordConfig.FooterIMG,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: p.Image,
		},
	}
	return &e
}


func (c *Controllers) GetImages(m *discordgo.MessageCreate) {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return
	}
	brand := info[1]
	pids, err := c.Repository.DB.GetPidByBrandName(brand)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	for _, v := range pids {
		emb := discordgo.MessageEmbed{
			Title:       v,
			Image:       &discordgo.MessageEmbedImage{
				URL:      fmt.Sprintf(fmt.Sprintf("https://imageresize.24i.com/?w=300&url=cache.mrporter.com/variants/images/%s/fr/w300.jpg", v)),
			},
		}
		_, err := c.Session.ChannelMessageSendEmbed(m.ChannelID, &emb)
		if err != nil {
			c.Logger.Error(err)
			c.BadAction(err.Error(), m)
			return
		}
		time.Sleep(3 * time.Second)
	}
	return
}
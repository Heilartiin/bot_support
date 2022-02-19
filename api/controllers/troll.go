package controllers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func (c *Controllers) Troll(m *discordgo.MessageCreate) {
	info := strings.Split(m.Content, " ")
	if len(info) < 3 {
		c.Logger.Error("Missing parameter")
		c.BadAction("Missing parameter", m)
		return
	}
	channelID, err := parser(info[1])
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	userName, err := parser(info[2])
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	colLink := discordgo.MessageEmbedField{
		Name: "Collection Links",
		Value: fmt.Sprintf(
			"[EtherScan](https://etherscan.io/address/0x3b3aa69d83b4102a1aca47344a174d8a8e9faf42?source=moby.gg)"),
		Inline: true,
	}
	txLink := discordgo.MessageEmbedField{
		Name: "Transaction Links",
		Value: fmt.Sprintf(
			"[EtherScan](https://etherscan.io/tx/0x1a4d368b8a84d8b2265fb333fecdfa0be3e51b3b7192e566d57f0744e660b574?source=moby.gg)"),
		Inline: true,
	}
	e := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			URL:          "https://moby.gg?utm_source=discord&utm_medium=webhook&utm_campaign=wallet_alerts&utm_content=footer",
			Name:         "Moby Alerts",
		},
		Title: "Unknown",
		URL:  "https://moby.gg/collection/0x3b3aa69d83b4102a1aca47344a174d8a8e9faf42?utm_source=discord&utm_medium=webhook&utm_campaign=wallet_alerts&utm_content=title",
		Description: fmt.Sprintf("%s minted Unknown — [view charts and stats on Moby](https://moby.gg/collection/0x3b3aa69d83b4102a1aca47344a174d8a8e9faf42?utm_source=discord&utm_medium=webhook&utm_campaign=wallet_alerts&utm_content=charts_and_stats).", userName),
		Fields: []*discordgo.MessageEmbedField{&colLink, &txLink},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      "https://cdn.discordapp.com/app-icons/885823483467161611/c17801ba641b336a0ef5b4e272ed3b3f.png?size=512",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("moby.gg"),
		},
		// 2015-12-31T12:00:00.000Z"
		Timestamp: fmt.Sprintf("%s", time.Now().Format("2006-01-02T15:04:05-0700")),
	}
	_, err = c.Session.ChannelMessageSendEmbed(channelID, &e)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	return
}

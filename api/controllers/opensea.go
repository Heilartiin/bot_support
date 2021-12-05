package controllers

import (
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (c *Controllers) OSGetCollectionInfo(m *discordgo.MessageCreate) {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error("Missing parameter")
		c.BadAction("Missing parameter", m)
		return
	}
	contract := info[1]
	res, err := c.OpenSea.GetInformationByContract(contract)
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	_, err = c.Session.ChannelMessageSendEmbed(m.ChannelID, c.createEmbedCollection(res))
	if err != nil {
		c.Logger.Error(err)
		c.BadAction(err.Error(), m)
		return
	}
	return
}

func (c *Controllers) createEmbedCollection(cc *models.OpenSeaCollection) *discordgo.MessageEmbed {

	contract := discordgo.MessageEmbedField{
		Name: 	"Contract",
		Value:	 fmt.Sprintf("[%s](%s)", cc.Address, cc.EtherscanUrl),
		Inline:  false,
	}

	txs := discordgo.MessageEmbedField{
		Name: 	"Txs",
		Value:	 fmt.Sprintf("[All Txs](%s)\n[Pending Txs](%s)",
			cc.TxsEtherscan, cc.PendingTxsEtherscan),
		Inline:  false,
	}

	floor := discordgo.MessageEmbedField{
		Name: "Floor",
		Value: fmt.Sprintf("ETH %.3f", cc.FloorPrice),
		Inline: true,
	}

	volume := discordgo.MessageEmbedField{
		Name: "Volume",
		Value: fmt.Sprintf("ETH %.3f", cc.TotalVolume),
		Inline: true,
	}

	owners := discordgo.MessageEmbedField{
		Name: "Owners",
		Value: fmt.Sprintf("%d", cc.NumOwners),
		Inline: true,
	}

	fees := discordgo.MessageEmbedField{
		Name: "Fees",
		Value: fmt.Sprintf("Service Fee %.1f %% +\nCreator Royalty %.1f %% =\nTotal Fee %.1f %%",
			cc.ServiceFee, cc.CreatorFee, cc.ServiceFee + cc.CreatorFee),
		Inline: false,
	}

	links := discordgo.MessageEmbedField{
		Name:   "Useful Links",
		Value: 	createLinks(cc),
		Inline: false,
	}


	fields := []*discordgo.MessageEmbedField{&contract, &txs, &floor, &volume, &owners, &fees, &links}

	fields = append(fields)

	e := discordgo.MessageEmbed{
		Title: "Collection " + cc.Name,
		URL:    cc.OSUrl,
		Fields: fields,
		Color:  45300,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    fmt.Sprintf("Contract Created %s | OS Created %s",
				cc.ContractCreated.Format("2006-01-02 15:04:05"),
				cc.OSCollectionCreated.Format("2006-01-02 15:04:05")),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:   cc.ImageUrl,
		},
	}
	return &e
}

func createLinks(res *models.OpenSeaCollection) string  {
	var links string
	if res.ExternalLink != "" {
		links += fmt.Sprintf("[Website](%s)\n", res.ExternalLink)
	}
	if res.DiscordUrl != "" {
		links += fmt.Sprintf("[Discord](%s)\n", res.DiscordUrl)
	}
	if res.TwitterUrl != "" {
		links += fmt.Sprintf("[Twitter](%s)\n", res.TwitterUrl)
	}
	if res.InstagramUrl != "" {
		links += fmt.Sprintf("[Instagram](%s)\n", res.InstagramUrl)
	}
	if res.TelegramUrl != "" {
		links += fmt.Sprintf("[Telegram](%s)\n", res.TelegramUrl)
	}
	return links
}

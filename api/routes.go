package api

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

//667588661470035989
func (api *API) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate)  {

	switch {
	// fix
	case m.ChannelID == "917449671323029505" && m.Author.ID != "917500852577648700" && s.State.User.ID != m.Author.ID:
		s.ChannelMessageSend(m.ChannelID, "discord.gg/" + m.Content)
		return
	// glory boyz
	case m.ChannelID == "917738838233870397" && m.Author.ID != "580517285085839381" && s.State.User.ID != m.Author.ID:
		s.ChannelMessageSend(m.ChannelID, "discord.gg/" + m.Content)
		return
	case !strings.HasPrefix(m.Content, api.Config.Discord.Prefix):
		return
	case m.Author.ID == s.State.User.ID:
		return
	}

	// Split command without prefix
	content := strings.Replace(m.Content, api.Config.Discord.Prefix, "", 1)
	command := strings.Fields(content)

	if len(command) == 0 {
		return
	}

	switch command[0] {
	case "os":
		api.Controllers.OSGetCollectionInfo(m)
	case "tx":
		api.Controllers.OSGetCollectionInfoByHash(m)
	case "troll":
		api.Controllers.Troll(m)
	//case "help":
	//	api.Controllers.Help(m)
	//case "atw":
	//	api.Controllers.GetAtw(m)
	//case "mrp":
	//	api.Controllers.GetImages(m)
	//case "atw-scraper":
	//	api.Controllers.GetATWFromScraper(m)
	//case "qt":
	//	api.Controllers.GetQTSToPrivateChannel(m)
	//case "qt-channel":
	//	api.Controllers.GetQTs(m)
	//case "vd-balance":
	//	api.Controllers.VDSinGetBalance(m)
	//case "vd-servers":
	//	api.Controllers.VDSinGetAllServers(m)
	//case "vd-delete":
	//	api.Controllers.VDSinDeleteAllServers(m)
	//case "1c-balance":
	//	api.Controllers.OneCloudGetBalance(m)
	//case "1c-servers":
	//	api.Controllers.OneCloudGetAllServers(m)
	//case "pm-proxies":
	//	api.Controllers.GetProxyMarketProxiesJSONFile(m)
	//case "pm-proxies-string":
	//	api.Controllers.GetProxyMarketStringJSONFile(m)
	//case "nike-storage":
	//	api.Controllers.StorageNikeEntries(m)
	//case "nike-tops":
	//	api.Controllers.GetEntriesByTOP(m)
	//case "nike-storage-accounts":
	//	api.Controllers.StorageNikeAccounts(m)
	//case "nike-accs":
	//	api.Controllers.GetNikeAccounts(m)
	//case "nike-bots":
	//	api.Controllers.GetTasksByAccounts(m)
	//case "nap-invisible":
	//	api.Controllers.GetNapProductsInvisible(m)
	default:
		return
		}
}

var TriggerWords = []string{"minted", "Popular NFT mint"}

func (api *API) parseMessage(s *discordgo.Session, m *discordgo.Message) bool {
	// Ignore all message prefix != "!"

	trigger := haveTrigger(m.Content)
	switch {
	case !trigger:
		return false
	case !strings.HasPrefix(m.Content, api.Config.Discord.Prefix):
		return false
	case m.Author.ID == s.State.User.ID:
		return false
	}

	if trigger {
		switch {
		case strings.Contains(m.Content, "Popular NFT mint"):

		}
	}
	// Split command without prefix
	content := strings.Replace(m.Content, api.Config.Discord.Prefix, "", 1)
	command := strings.Fields(content)

	if len(command) == 0 {
		return false
	}
	return true
}

func haveTrigger(content string) bool {
	for _, v := range TriggerWords {
		if strings.Contains(content, v) {
			return true
		}
	}
	return false
}

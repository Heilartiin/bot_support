package api

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

//667588661470035989
func (api *API) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate)  {

	// Ignore all message prefix != "!"
	if !strings.HasPrefix(m.Content, api.Config.Discord.Prefix) {
		return
	}
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Split command without prefix
	content := strings.Replace(m.Content, api.Config.Discord.Prefix, "", 1)
	command := strings.Fields(content)

	if len(command) == 0 {
		return
	}

	switch command[0] {
	case "help":
		api.Controllers.Help(m)
	case "atw":
		api.Controllers.GetAtw(m)
	case "mrp":
		api.Controllers.GetImages(m)
	case "qt":
		api.Controllers.GetQTSToPrivateChannel(m)
	case "qt-channel":
		api.Controllers.GetQTs(m)
	case "vd-balance":
		api.Controllers.VDSinGetBalance(m)
	case "vd-servers":
		api.Controllers.VDSinGetAllServers(m)
	case "vd-delete":
		api.Controllers.VDSinDeleteAllServers(m)
	case "1c-balance":
		api.Controllers.OneCloudGetBalance(m)
	case "1c-servers":
		api.Controllers.OneCloudGetAllServers(m)
	case "pm-proxies":
		api.Controllers.GetProxyMarketProxiesJSONFile(m)
	case "pm-proxies-string":
		api.Controllers.GetProxyMarketStringJSONFile(m)
	case "nike-storage":
		api.Controllers.StorageNikeEntries(m)
	case "nike-tops":
		api.Controllers.GetEntriesByTOP(m)
		default:
			return
		}
}



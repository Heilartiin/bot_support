package controllers

import (
	"github.com/bwmarrin/discordgo"
)

func (c *Controllers) GetTasksByAccounts(m *discordgo.MessageCreate)  {
	//info := strings.Split(m.Content, " ")
	//if len(info) < 5 {
	//	c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
	//	c.BadAction("Missing parameter", m)
	//	return
	//}
	//productID := info[1]
	//launchID  := info[2]
	//split     := info[3]
	//ValidSizes     := info[4]
	//
	//splitNumber, err := strconv.Atoi(split)
	//if err != nil {
	//	c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
	//	c.BadAction(err.Error(), m)
	//	return
	//}


	//tasks, err := c.Repository.DB.GetAccountAllActive()
	//if err != nil {
	//	c.Logger.Error(errors.WithStack(err))
	//	c.BadAction(err.Error(), m)
	//	return
	//}

	//f := func(i int) bool {
	//	if i > len(tasks) {
	//		return true
	//	}
	//	return false
	//}

	//var tasksSlices [][]*models.NikeBotTask
	//for i := 0; i < len(tasks); i += splitNumber {
	//	for {
	//		if f(i) {
	//			i -= 1
	//			break
	//		}
	//	}
	//	fmt.Println()
	//}

	//
	//channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	//if err != nil {
	//	c.Logger.Error(errors.WithStack(err))
	//	c.BadAction(err.Error(), m)
	//	return
	//}
	//config := models.BotConfig{
	//	DelayMin:           0,
	//	DelayMax:           0,
	//	SubmitRetry:        0,
	//	ReleaseTime:        time.Time{},
	//	PrepareSessionTime: time.Time{},
	//	EntryTime:          time.Time{},
	//	Tasks:              tasks,
	//}
	//f, err := json.Marshal(config)
	//if err != nil {
	//	c.Logger.Error(errors.WithStack(err))
	//	c.BadAction(err.Error(), m)
	//	return
	//}
	//var prettyJSON bytes.Buffer
	//err = json.Indent(&prettyJSON, f,"", "\t")
	//if err != nil {
	//	c.Logger.Error(errors.WithStack(err))
	//	c.BadAction(err.Error(), m)
	//	return
	//}
	//nameFile := fmt.Sprintf("nike_tasks-%s.json", time.Now().Format("15:04:05"))
	//_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, &prettyJSON)
	//if err != nil {
	//	c.Logger.Error(errors.WithStack(err))
	//	c.BadAction(err.Error(), m)
	//	return
	//}
	return
}

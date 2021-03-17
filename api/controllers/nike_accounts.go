package controllers

import (
	"bytes"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

func (c *Controllers) StorageNikeAccounts(m *discordgo.MessageCreate) {
	b, err := c.OpenFile(m)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(fmt.Sprintf("Error for open file: %s", err), m)
		return
	}
	var accounts []*models.NikeAccount
	err = gocsv.Unmarshal(bytes.NewBuffer(b), &accounts)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	var load int
	for _, v := range accounts {
		_, err := c.Repository.DB.CreateNikeAccount(v)
		if err != nil {
			if err.Error() =="pq: duplicate key value violates unique constraint \"nike_accounts_login_key\"" {
				continue
			}
			c.Logger.Error(errors.WithStack(err))
			c.BadAction(err.Error(), m)
			return
		}
		load += 1
	}
	c.SuccessAction(fmt.Sprintf("Number of accounts loaded: %d", load), m)
	return
}

func (c *Controllers) GetNikeAccounts(m *discordgo.MessageCreate)  {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter: priority id")))
		c.BadAction("Missing parameter: priority id", m)
		return
	}
	priorityID := info[1]
	priorityIDInt, err := strconv.Atoi(priorityID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(fmt.Sprintf("\n 1 - общие \n 2 - Артур \n 3 - Дима \n 4 - остальные"), m)
		return
	}
	var resp []*models.NikeAccount
	switch priorityIDInt {
	case 5:
		resp, err = c.Repository.DB.GetAllAccounts()
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			c.BadAction(err.Error(), m)
			return
		}
	case 1, 2, 3, 4:
		resp, err = c.Repository.DB.GetAccountsByPriority(priorityIDInt)
		if err != nil {
			c.Logger.Error(errors.WithStack(err))
			c.BadAction(err.Error(), m)
			return
		}
	default:
		c.Logger.Error(errors.WithStack(err))
		c.BadAction("unknown priority id: " + strconv.Itoa(priorityIDInt), m)
		return
	}

	csvContent, err := gocsv.MarshalString(&resp)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
	br := []byte(csvContent)
	nameFile := fmt.Sprintf("nike-accounts-%s.csv", time.Now().Format("15:04:05"))
	_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, bytes.NewBuffer(br))
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}

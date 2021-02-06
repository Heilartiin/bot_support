package controllers

import (
	"bytes"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (c *Controllers) doReq(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return resp, nil
}


func (c *Controllers) OpenFile(m *discordgo.MessageCreate) ([]byte, error) {
	if len(m.Attachments) == 0 {
		c.BadAction("No files found", m)
		return nil, errors.WithStack(errors.New("No files found"))
	}
	file := m.Attachments[0]
	resp, err := c.doReq(file.URL)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return b, nil
}

func (c *Controllers) StorageNikeEntries(m *discordgo.MessageCreate) {
	b, err := c.OpenFile(m)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(fmt.Sprintf("Error for open file: %s", err), m)
		return
	}
	var entries []*models.NikeEntry
	err = gocsv.Unmarshal(bytes.NewBuffer(b), &entries)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	var all int
	for _, v := range entries {
		_, err := c.Repository.DB.CreateNikeEntry(v)
		if err != nil {
			if err.Error() =="pq: duplicate key value violates unique constraint \"nike_entries_username_launch_key\"" {
				continue
			}
			c.Logger.Error(errors.WithStack(err))
			c.BadAction(err.Error(), m)
			return
		}
		all += 1
	}
	c.SuccessAction(fmt.Sprintf("Number of entries loaded: %d", all), m)
	return
}

func (c *Controllers) GetEntriesByTOP(m *discordgo.MessageCreate)  {
	info := strings.Split(m.Content, " ")
	if len(info) < 2 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter: Launch")))
		c.BadAction("Missing parameter: Launch", m)
		return
	}
	launch := info[1]
	resp, err := c.Repository.DB.GetNikeEntriesSortByTime(launch)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
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
	embedMessage, err := c.CreateEmbedForNikeTOP(resp)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	_, err = c.Session.ChannelMessageSendEmbed(channelPrivate.ID, embedMessage)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		c.BadAction(err.Error(), m)
		return
	}
	br := []byte(csvContent)
	nameFile := fmt.Sprintf("nike-tops-%s.csv", time.Now().Format("15:04:05"))
	_, err = c.Session.ChannelFileSend(channelPrivate.ID, nameFile, bytes.NewBuffer(br))
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return
	}
}

func (c *Controllers) CreateEmbedForNikeTOP(entries []*models.NikeEntry) (*discordgo.MessageEmbed, error)  {
	var completed []*models.NikeEntry
	for index, v := range entries {
		if index > 10 {
			continue
		}
		completed = append(completed, v)
	}
	if len(entries) == 0 {
		return nil, errors.New("Entries is empty")
	}
	var usernames []string
	var times  []string
	var statuses []string
	for _, v := range completed {
		usernames = append(usernames, fmt.Sprintf("> %s",v.Username))
		times = append(times, v.EntryTime.Format("15:04:05.000Z"))
		statuses = append(statuses, v.Status)
	}
	fieldUserName := discordgo.MessageEmbedField{
		Name:   "UserName",
		Value:  strings.Join(usernames[:],"\n"),
		Inline: true,
	}
	fieldTimes := discordgo.MessageEmbedField{
		Name:   "Entry Time",
		Value:  strings.Join(times[:],"\n"),
		Inline: true,
	}
	fieldStatuses := discordgo.MessageEmbedField{
		Name:   "Status",
		Value:  strings.Join(statuses[:],"\n"),
		Inline: true,
	}
	ent := entries[0]
	var fields []*discordgo.MessageEmbedField
	fields = append(fields, &fieldUserName, &fieldTimes, &fieldStatuses)
	dm := discordgo.MessageEmbed{
		Title:       "Nike TOP ENTRIES",
		Author:      &discordgo.MessageEmbedAuthor{
			URL:          ent.Launch,
			Name:         ent.Launch,
		},
		Fields:      fields,
	}
	return &dm, nil
}

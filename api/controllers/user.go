package controllers

import (
	"context"
	"github.com/Heilartin/bot_support/models"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
	"time"
)


func (c *Controllers) CreateUser(m *discordgo.MessageCreate) error  {
	info := strings.Split(m.Content, " ")
	if len(info) < 3 {
		c.Logger.Error(errors.WithStack(errors.New("Missing parameter")))
		c.BadAction("Missing parameter", m)
		return errors.New("Missing parameter")
	}
	channelPrivate, err := c.Session.UserChannelCreate(m.Author.ID)
	if err != nil {
		c.BadAction(err.Error(), m)
		return err
	}
	wishlist := info[1]
	token := info[2]
	u := models.User{
		MemberID: m.Author.ID,
		Token: token,
		WishList: wishlist,
		PrivateChannel: channelPrivate.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}
	userID, err := c.Repository.DB.CreateUser(u)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return  errors.WithStack(err)
	}
	u.ID = userID
	e := c.CreateUserEmbed("Information", &u)
	_, err = c.Session.ChannelMessageSendEmbed(u.PrivateChannel, e)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (c *Controllers) GetUser(ctx context.Context, m *discordgo.MessageCreate) error  {
	userInfo, err := c.UserFromContext(ctx)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	e := c.CreateUserEmbed("Information", userInfo)
	_, err = c.Session.ChannelMessageSendEmbed(userInfo.PrivateChannel, e)
	if err != nil {
		c.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (c *Controllers) UserFromContext(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value("userData").(models.User)
	if !ok {
		return nil, errors.WithStack(errors.New("No user data"))
	}
	return &user, nil
}
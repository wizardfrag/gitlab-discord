package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/wizardfrag/gitlab-discord/gitlabdiscord"
	"log"
)

type Bot struct {
	ID string
}

func Run(config gitlabdiscord.Config) (*Bot, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return nil, err
	}

	u, err := session.User("@me")
	if err != nil {
		log.Println("error obtaining account details,", err)
		return nil, err
	}

	bot := &Bot{
		ID: u.ID,
	}

	session.AddHandler(bot.messageCreate)

	err = session.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return nil, err
	}

	return bot, nil
}

func (b *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages that the bot sent!
	if m.Author.ID == b.ID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping!")
	}
}

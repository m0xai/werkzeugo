package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Discord struct {
	session *discordgo.Session
}

func New(token string) (*Discord, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	err = s.Open()
	if err != nil {
		return nil, err
	}
	return &Discord{session: s}, nil
}

func (d Discord) SetupHandlers() {
	d.session.AddHandler()
}

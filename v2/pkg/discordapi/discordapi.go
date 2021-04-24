package discordapi

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	instance *DiscordAPI
)

type DiscordAPI struct {
	session *discordgo.Session
}

func GetAPI() (*DiscordAPI, error) {
	if instance == nil {
		newAPI := &DiscordAPI{}
		discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_AUTH_TOKEN"))
		if err != nil {
			return nil, err
		}
		newAPI.session = discord
		instance = newAPI
	}
	return instance, nil
}

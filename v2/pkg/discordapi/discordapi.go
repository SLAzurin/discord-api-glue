package discordapi

import (
	"errors"
	"os"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
	"github.com/bwmarrin/discordgo"
)

var (
	instance                  *DiscordAPI
	listeningDiscordChannelID *string
	subscribers               map[string]*chan genericapi.APIMessage
	ListenChannel             *chan genericapi.APIMessage
)

// DiscordAPI is the struct of the DiscordAPI inside the API Glue app
type DiscordAPI struct {
	session       *discordgo.Session
	ListenChannel *chan genericapi.APIMessage
	channelOpen   bool
}

func GetAPI() (*DiscordAPI, error) {
	if instance == nil {
		// Setup new DiscordAPI
		newAPI := &DiscordAPI{}

		// Get Discord channel to listen (from .env)
		dChan := os.Getenv("DISCORD_CHANNEL")
		listeningDiscordChannelID = nil

		// Create the new discord instance
		discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_AUTH_TOKEN"))
		if err != nil {
			return nil, err
		}
		newAPI.session = discord
		tChan := make(chan genericapi.APIMessage)
		ListenChannel = &tChan
		newAPI.ListenChannel = ListenChannel
		newAPI.channelOpen = true

		// Register handlers
		newAPI.session.AddHandler(messageCreateHandler)

		//Register intents
		newAPI.session.Identify.Intents = discordgo.IntentsGuildMessages

		// Connect to discord
		err = newAPI.session.Open()
		if err != nil {
			return nil, err
		}

		// Get Discord channel ID
		if len(newAPI.session.State.Guilds) != 1 {
			return nil, errors.New("You can only use this bot in 1 discord server")
		}
		channels, _ := newAPI.session.GuildChannels(newAPI.session.State.Guilds[0].ID)
		for _, ch := range channels {
			if ch.Type != discordgo.ChannelTypeGuildText {
				continue
			}
			listeningDiscordChannelID = &ch.ID
		}
		if listeningDiscordChannelID == nil {
			return nil, errors.New("Cannot listen to channel:" + dChan)
		}

		instance = newAPI
		go listenToGoChannel()
	}
	return instance, nil
}

// This publishes the incoming messages to the subscribers
func listenToGoChannel() {
	for {
		if instance.channelOpen {
			for elem := range *instance.ListenChannel {
				// Incoming messages to send to discord
				if len(elem.Author) != 0 {
					elem.Author += ": "
				}
				instance.session.ChannelMessageSend(*listeningDiscordChannelID, elem.Author+elem.Content)
			}
		}
	}
}

func (*DiscordAPI) Subscribe(name string, c *chan genericapi.APIMessage) error {
	if _, ok := subscribers[name]; ok {
		return errors.New("Subscriber already exists")
	}
	subscribers[name] = c
	return nil
}

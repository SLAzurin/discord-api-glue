package discordapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
	"github.com/bwmarrin/discordgo"
)

var (
	instance                 *DiscordAPI
	listeningDiscordChannels map[string]*string
	subscribers              map[string]*chan genericapi.APIMessage
	ListenChannel            *chan genericapi.APIMessage
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

		// Get list of Discord channels to listen (from .env)
		listeningDiscordChannels = make(map[string]*string)
		jsonArr := []string{}
		err := json.Unmarshal([]byte(os.Getenv("DISCORD_CHANNELS")), &jsonArr)
		if err != nil {
			return nil, errors.New("DISCORD_CHANNELS not JSON!")
		}
		for _, dChan := range jsonArr {
			listeningDiscordChannels[dChan] = nil
		}

		// Finally create the new instance
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
			if _, ok := listeningDiscordChannels[ch.Name]; ok {
				listeningDiscordChannels[ch.Name] = &ch.ID
				fmt.Println("Listening Channel:", ch.Name)
			} else {
				fmt.Println("Ignoring Channel:", ch.Name)
			}
		}
		// register handlers

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
				// This is a testing line
				// This will be replaced with sending the actual message to discord
				fmt.Println("Discord[", elem.Destination, "]:", elem.Content)
				for _, subbersChan := range subscribers {
					*subbersChan <- elem
				}
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

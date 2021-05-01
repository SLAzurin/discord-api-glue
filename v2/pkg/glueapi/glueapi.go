package glueapi

import (
	"fmt"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/discordapi"
	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
)

var (
	discordAPI *discordapi.DiscordAPI
	// Add chat api here too
	incomingDiscordMessages *chan genericapi.APIMessage
	instance                *GlueAPI
)

// GlueAPI
type GlueAPI struct {
}

const (
	TYPE_DISCORD = "DISCORD"
	TYPE_GCHAT   = "GCHAT"
)

type GlueAPIMessage struct {
	DestinationPlatform string
	Payload             genericapi.APIMessage
}

// Start is the entrypoint function
func GetAPI() (*GlueAPI, error) {
	if instance == nil {
		api := GlueAPI{}
		dapi, err := discordapi.GetAPI()
		if err != nil {
			return nil, err
		}
		discordAPI = dapi
		t := make(chan genericapi.APIMessage)
		incomingDiscordMessages = &t
		discordAPI.Subscribe("GlueAPI", incomingDiscordMessages)
		instance = &api

		go listenToIncomingDiscordMessages()
	}
	return instance, nil
}

func listenToIncomingDiscordMessages() {
	for {
		for elem := range *incomingDiscordMessages {
			// Incoming messages to send to GChat
			if len(elem.Author) != 0 {
				elem.Author += ": "
			}
			// Send to GChat
			fmt.Println(elem.Author + elem.Content)
		}
	}
}

package glueapi

import (
	"github.com/SLAzurin/discord-api-glue/v2/pkg/discordapi"
	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
)

var (
	discordAPI *discordapi.DiscordAPI
	// Add chat api here too
	instance *GlueAPI
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
		instance = &api
	}
	return instance, nil
}

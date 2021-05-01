package discordapi

import (
	"errors"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
)

func (*DiscordAPI) Subscribe(name string, c *chan genericapi.APIMessage) error {
	if _, ok := subscribers[name]; ok {
		return errors.New("Subscriber already exists")
	}
	subscribers[name] = c
	return nil
}

func publishMessage(m *genericapi.APIMessage) {
	for _, sub := range subscribers {
		*sub <- *m
	}
}
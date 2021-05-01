package discordapi

import (
	"log"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
	"github.com/bwmarrin/discordgo"
)

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	incomingAPIMessage := genericapi.APIMessage{}
	content, err := contentWithMoreMentionsReplaced(m.Message, instance.session)
	if err != nil {
		log.Println(err)
	}

	incomingAPIMessage.Content = content

	if m.Member.Nick != "" {
		incomingAPIMessage.Author = m.Member.Nick
	} else {
		incomingAPIMessage.Author = m.Author.Username
	}
	// fmt.Println(incomingAPIMessage.Author + ": " + incomingAPIMessage.Content)
	go publishMessage(&incomingAPIMessage)
}

package discordapi

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(m.Author.Username+":", m.Content)
}

package discordapi

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func contentWithMoreMentionsReplaced(m *discordgo.Message, s *discordgo.Session) (content string, err error) {
	content = m.Content

	for _, user := range m.Mentions {
		nick := user.Username

		member, err := s.GuildMember(*guildID, user.ID)
		if err == nil && member.Nick != "" {
			nick = member.Nick
		}

		content = strings.NewReplacer(
			"<@"+user.ID+">", "@"+user.Username,
			"<@!"+user.ID+">", "@"+nick,
		).Replace(content)
	}

	// Replace Roles
	guildStruct, err := s.Guild(*guildID)
	if err != nil {
		log.Println("Getting Guild error", err)
	}
	roles := guildStruct.Roles
	patternRoles := regexp.MustCompile("<@&[^>]*>")
	content = patternRoles.ReplaceAllStringFunc(content, func(mention string) string {
		role := mention[3 : len(mention)-1]
		for _, roleStruct := range roles {
			if roleStruct.ID == role {
				return "@" + roleStruct.Name
			}
		}
		return "@unknown-role"
	})

	// Replace Channels
	var patternChannels = regexp.MustCompile("<#[^>]*>")
	content = patternChannels.ReplaceAllStringFunc(content, func(mention string) string {
		channel, err := s.State.Channel(mention[2 : len(mention)-1])
		if err != nil || channel.Type == discordgo.ChannelTypeGuildVoice {
			return mention
		}

		return "#" + channel.Name
	})
	return
}

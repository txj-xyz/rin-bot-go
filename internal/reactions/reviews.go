package reactions

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ReviewReaction(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("[INFO] Reacted to: %v, Message URL: https://discord.com/channels/%v/%v/%v\n",
		m.Message.Author.Username,
		m.Message.GuildID,
		m.Message.ChannelID,
		m.Message.ID)

	// Send out the reactions here
	s.MessageReactionAdd(m.ChannelID, m.Message.ID, "🥳")
	s.MessageReactionAdd(m.ChannelID, m.Message.ID, "💜")
	s.MessageReactionAdd(m.ChannelID, m.Message.ID, "🎉")
}

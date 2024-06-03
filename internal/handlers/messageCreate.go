package handlers

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/txj-xyz/rin-bot-go/internal/reactions"
)

var (
	ChannelReactListen string
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Load up the reaction channel here
	ChannelReactListen = os.Getenv("MESSAGE_REACT_CHANNEL")
	if ChannelReactListen == "" {
		log.Fatal("[ERROR] Could not find reaction channel ID from .env please try again")
	}

	// Ignore bot owned messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the author of the message is a bot then we ignore it
	if m.Author.Bot == true {
		return
	}

	if m.ChannelID == ChannelReactListen {
		reactions.ReviewReaction(s, m)
	}

}

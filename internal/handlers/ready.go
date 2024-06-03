package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Ready(_ *discordgo.Session, r *discordgo.Ready) {
	log.Printf("[INFO] Logged in as %s", r.User.String())
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Setup the token variable here
var (
	Token              string
	ChannelReactListen string
)

// pre main hook start check flags
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR] Loading .env file please create one and load again. Exiting")
	}

	// Load token here
	Token = os.Getenv("TOKEN")
	if Token == "" {
		log.Fatal("[ERROR] Loading Discord Bot 'TOKEN' please try again. Exiting")
	}

	// Load up the reaction channel here
	ChannelReactListen = os.Getenv("MESSAGE_REACT_CHANNEL")
	if ChannelReactListen == "" {
		log.Fatal("[ERROR] Could not find reaction channel ID from .env please try again")
	}
}

func main() {

	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("[ERROR] Failed to create new bot client using token, exiting")
		return
	}

	// Register the message create event to the bot client
	discord.AddHandler(messageCreate)

	// gracefully handle control-C exits
	discord.AddHandler(func(_ *discordgo.Session, r *discordgo.Ready) {
		log.Printf("[INFO] Logged in as %s", r.User.String())
	})

	// Register the flags needed for these perms
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// open websocket here to the gateway
	err = discord.Open()
	if err != nil {
		log.Fatalf("[ERROR] Opening connection to discord, %s", err)
	}

	// Create our signal listen notifier so when we attempt to escape we send it to the required
	// Channel to exit properly instead of crashing fatally
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// close when done
	err = discord.Close()
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to the discord gateway possibly ratelimited, please try again.")
	}

}

// Message creation hook to listen to messages
// Skip all bot based messages and only listen to the ChannelID from the .env file
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore bot owned messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the author of the message is a bot then we ignore it
	if m.Author.Bot == true {
		return
	}

	if m.ChannelID == ChannelReactListen {
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
}

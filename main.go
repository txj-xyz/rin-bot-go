package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/txj-xyz/rin-bot-go/events"
)

// Setup the token variable here
var (
	Token string
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
}

// Main bot thread
func main() {

	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("[ERROR] Failed to create new bot client using token, exiting")
		return
	}

	// Register the message create event to the bot client
	discord.AddHandler(events.MessageCreate)

	// gracefully handle control-C exits
	discord.AddHandler(events.Ready)

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

package bot

import (
	"discordbot/internal/models/crawler"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	OpenWeatherToken string
	BotToken         string
)

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Add event handler
	discord.AddHandler(newMessage)

	// Open session
	discord.Open()
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	log.Println(message.Content)

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Contains(message.Content, "weather"):
		discord.ChannelMessageSend(message.ChannelID, "I can help with that!")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	case strings.Contains(message.Content, "maimai"):
		discord.ChannelMessageSend(message.ChannelID, "幹 你戳到測試了 操 我還在搞Golang的機器人")
		crawler.Run("8058430092794")
		// discord.ChannelMessageSend(message.ChannelID, crawler.Run("8058430092794"))
	}
}

package bot

import (
	"bytes"
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
	SID              string
	Password         string
)

func Run(sid string, pwd string) {

	// Create new Discord Session
	discord, err := discordgo.New("Bot " + BotToken)
	SID = sid
	Password = pwd
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
	// log.Println(message.Content)

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, "&maimai") {
		args := strings.Fields(message.Content)

		if len(args) < 2 {
			discord.ChannelMessageSend(message.ChannelID, "請提供好友ID.")
			return
		}

		id := args[1]
		// log.Println(id)
		pic := crawler.Run(id, SID, Password)
		reader := bytes.NewReader(pic)
		discord.ChannelFileSendWithMessage(message.ChannelID, "", "maimai.png", reader)
	}
}

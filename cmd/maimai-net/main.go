package main

import (
	"discordbot/internal/models/bot"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Server started")
	sid := os.Getenv("SEGA_ID")
	pwd := os.Getenv("SEGA_PW")
	botToken := os.Getenv("DISCORD_TOKEN")

	// // Start the bot
	bot.BotToken = botToken
	bot.Run(sid, pwd)

	// for {
	// 	input := []byte("friendid")
	// 	crawler.Run(string(input), sid, pwd)
	// }
}

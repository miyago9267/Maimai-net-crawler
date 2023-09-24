package main

import (
	"fmt"
	"os"

	"discordbot/internal/models/bot"
	"discordbot/internal/models/crawler"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Hello, World!")
	botToken := os.Getenv("DISCORD_TOKEN")

	// Start the bot
	bot.BotToken = botToken
	crawler.Run("8058430092794")
	bot.Run()
}

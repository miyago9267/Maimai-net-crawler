package main

import (
	"bufio"
	"log"
	"maimainet-crawler/internal/crawler"
	"os"

	"maimainet-crawler/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.InitConfig()

	sid := config.Cfg.SEGA_ID
	pwd := config.Cfg.SEGA_PW

	// // Start the bot

	id := bufio.NewReader(os.Stdin)

	input, err := id.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	crawler.Run(input, sid, pwd)

}

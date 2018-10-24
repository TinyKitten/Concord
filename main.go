package main

import (
	"log"
	"os"

	"github.com/TinyKitten/Concord/discord"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d, err := discord.NewDiscord(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))

	emojis, err := api.GetEmoji()
	if err != nil {
		log.Fatal(err)
	}

	gID := os.Getenv("DISCORD_GUILD_ID")
	for key, url := range emojis {
		err = d.TransferEmoji(gID, key, url)
		if err != nil {
			log.Fatalf("An error occurred during processing %s: %s", key, err.Error())
		}
		log.Printf("Uploaded: %s", key)
	}
}

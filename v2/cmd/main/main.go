package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/discordapi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Initializing Discord Glue...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// TODO: Replace discord by GlueAPI
	discordAPI, err := discordapi.GetAPI()
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}

	// The code below is for testing messaging
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		*discordAPI.ListenChannel <- discordapi.DiscordAPIMessage{Content: text, Destination: "general", Author: "THIS_BOT"}
	}
}

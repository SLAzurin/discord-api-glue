package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/SLAzurin/discord-api-glue/v2/pkg/genericapi"
	"github.com/SLAzurin/discord-api-glue/v2/pkg/glueapi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Initializing Discord Glue...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// TODO: Replace discord by GlueAPI
	glueAPI, err := glueapi.GetAPI()
	if err != nil {
		log.Fatalln("error creating GlueAPI,", err)
	}
	
	// The code below is for testing messaging
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		go glueAPI.SendMessage(
			glueapi.GlueAPIMessage{
				DestinationPlatform: glueapi.TYPE_DISCORD,
				Payload: genericapi.APIMessage{
					Content:     text,
					Destination: "general",
					Author:      "THIS_BOT",
				},
			},
		)
	}
}

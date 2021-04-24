package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/SLAzurin/discord-api-glue/v2/pkg/discordapi"
)

func main() {
	fmt.Println("Initializing Discord Glue...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	_, err = discordapi.GetAPI()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

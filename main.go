package main

import (
	"context"
	"fmt"
	"log"

	"github.com/A-K-2001/slack-test/database"
	"github.com/A-K-2001/slack-test/internal/config"
	"github.com/A-K-2001/slack-test/internal/linear"
	"github.com/A-K-2001/slack-test/internal/slackfunc"
	"github.com/A-K-2001/slack-test/internal/store"
)

func main() {

	config, err := config.LoadConfig()

	//database connect

	postgres := store.NewPostgresStorage()

	if err := postgres.Connect(context.Background(), config.DATABASE_URL); err != nil {
		fmt.Println("Error connecting to database", err)
	}

	defer postgres.Close()
	repo := database.NewRepository(postgres.Pool)

	client := slackfunc.NewClient(config.SLACK_TOKEN, config.BOT_TOKEN)
	linear.CreateLinearClient(config.Linear_api_key)
	socketClient := slackfunc.NewSocketClient(client)
	go slackfunc.ListenEvnet(socketClient, client, repo)

	err = socketClient.Run()
	if err != nil {
		log.Fatal(err)
	}
}

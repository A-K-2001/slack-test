package slackfunc

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func NewClient(appToken, botToken string) *slack.Client {
	return slack.New(
		botToken,
		slack.OptionAppLevelToken(appToken),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)
}

func NewSocketClient(client *slack.Client) *socketmode.Client {
	return socketmode.New(
		client,
		socketmode.OptionDebug(true),
		// socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
		socketmode.OptionDebug(false),
	)
}

package slackfunc

import (
	"fmt"
	"strings"

	"github.com/A-K-2001/slack-test/database"
	"github.com/A-K-2001/slack-test/internal/linear"
	"github.com/A-K-2001/slack-test/internal/slackcmd"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func ListenEvnet(socketClient *socketmode.Client, client *slack.Client, repo *database.Repository) {
	for evt := range socketClient.Events {
		switch evt.Type {
		case socketmode.EventTypeInteractive:
			interactionEvent, ok := evt.Data.(slack.InteractionCallback)
			if !ok {
				fmt.Printf("Ignored %+v\n", evt)
				continue
			}
			socketClient.Ack(*evt.Request)
			handleInteraction(interactionEvent, client)
		case socketmode.EventTypeSlashCommand:
			cmd, ok := evt.Data.(slack.SlashCommand)
			if !ok {
				fmt.Printf("Ignored %+v\n", evt)
				continue
			}
			socketClient.Ack(*evt.Request)
			handleSlashCommand(cmd, client)
		case socketmode.EventTypeEventsAPI:
			eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
			if !ok {
				fmt.Printf("Ignored %+v\n", evt)
				continue
			}

			socketClient.Ack(*evt.Request)

			switch eventsAPIEvent.Type {
			case slackevents.CallbackEvent:
				innerEvent := eventsAPIEvent.InnerEvent
				switch ev := innerEvent.Data.(type) {
				case *slackevents.FileSharedEvent:
					handleFileShared(ev, client)
				case *slackevents.MessageEvent:
					handleMessage(ev, client, repo)
				}
			}
		}
	}
}

func handleSlashCommand(cmd slack.SlashCommand, client *slack.Client) {
	switch cmd.Command {
	case "/issues":
		slackcmd.HandleIssuesCommand(cmd, client)
	default:
		_, _, err := client.PostMessage(cmd.ChannelID, slack.MsgOptionText("Command not recognized", false))
		if err != nil {
			fmt.Printf("Error responding to command: %v\n", err)
		}
	}
}

func handleFileShared(ev *slackevents.FileSharedEvent, client *slack.Client) {
	// channel_id := ev.ChannelID
	msg := linear.GetIssueId(ev, client)
	msgParts := strings.Split(msg, ":")
	issueID := ""
	if len(msgParts) > 1 {
		issueID = msgParts[1]
		// Use issueID as needed
	}

	fmt.Println("issueId: ", issueID)
	file, _, _, err := client.GetFileInfo(ev.FileID, 0, 0)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Printf("File shared: %s\n", file.Name)

	err = linear.UploadScreenshotToIssue(issueID, file.URLPrivate, file.Name)
	if err != nil {
		fmt.Printf("Error upload img on linear: %v\n", err)
		return
	}

}

func handleInteraction(callback slack.InteractionCallback, client *slack.Client) {
	switch callback.Type {
	case slack.InteractionTypeViewSubmission:
		viewState := callback.View.State.Values

		title := viewState["title_block"]["title_input"].Value
		description := viewState["description_block"]["description_input"].Value
		team := viewState["team_block"]["team_input"].SelectedOption.Value
		vale := viewState["team_block"]["team_input"].SelectedOption.Text.Text

		fmt.Printf("Form submitted.\nTitle: %s\nDescription: %s\nTeam: %s : %s", title, description, vale, team)

		issueID, err := linear.CreateTriage(team, title, description)

		if err != nil {
			fmt.Printf("Error creating issue on linear: %v\n", err)
		}
		message := "Thanks for submitting the issue! Please upload screenshot to support your issue in the thread of this chat. \nYour issueId: " + issueID + "\nTitle: " + title + "\nDescription: " + description

		_, _, err = client.PostMessage(callback.User.ID, slack.MsgOptionText(message, false))
		if err != nil {
			fmt.Printf("Error sending upload instructions: %v\n", err)
		}
	case slack.InteractionTypeViewClosed:
		fmt.Print("Input form closed.\n")
	}
}

func handleMessage(ev *slackevents.MessageEvent, client *slack.Client, repo *database.Repository) {
	print(ev.Text)
	switch strings.ToLower(strings.Split(ev.Text, " ")[0]) {
	case "rider":
		slackcmd.AppCammonds(ev, client, repo)
	default:
		client.PostMessage(ev.Channel, slack.MsgOptionText("comming soon ...", true))
	}
}

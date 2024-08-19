package slackcmd

import (
	"fmt"

	"github.com/slack-go/slack"
)

func HandleIssuesCommand(cmd slack.SlashCommand, client *slack.Client) {
	modalRequest := slack.ModalViewRequest{
		Type:       slack.ViewType("modal"),
		CallbackID: "input-modal-callback-id",
		Title: &slack.TextBlockObject{
			Type: "plain_text",
			Text: "Report issue",
		},
		Submit: &slack.TextBlockObject{
			Type: "plain_text",
			Text: "Submit",
		},
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				&slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: "title_block",
					Label: &slack.TextBlockObject{
						Type: "plain_text",
						Text: "Title",
					},
					Element: &slack.PlainTextInputBlockElement{
						Type:     slack.METPlainTextInput,
						ActionID: "title_input",
					},
				},
				&slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: "description_block",
					Label: &slack.TextBlockObject{
						Type: "plain_text",
						Text: "Description",
					},
					Element: &slack.PlainTextInputBlockElement{
						Type:      slack.METPlainTextInput,
						ActionID:  "description_input",
						Multiline: true,
					},
				},
				&slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: "team_block",
					Label: &slack.TextBlockObject{
						Type: "plain_text",
						Text: "Team",
					},
					Element: &slack.SelectBlockElement{
						Type:     slack.OptTypeStatic,
						ActionID: "team_input",
						Options: []*slack.OptionBlockObject{
							{
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: "Flow",
								},
								Value: "7ed04cf1-34c3-42d1-80cd-86768a49579c",
							},
							{
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: "Infrastructure",
								},
								Value: "b9666332-5eb9-48a8-a7b1-4bf0014ebd25",
							},
							{
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: "API",
								},
								Value: "5e8be61e-1600-44ae-bdf5-a160393c18e0",
							},
							{
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: "UserMobile",
								},
								Value: "4dae451f-25cf-40ad-acfe-49a68a128ba1",
							},
							{
								Text: &slack.TextBlockObject{
									Type: "plain_text",
									Text: "Tech",
								},
								Value: "a7020cf4-5c82-4da9-9034-1b2dbfd3c78c",
							},
						},
					},
				},
			},
		},
	}

	_, err := client.OpenView(cmd.TriggerID, modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %v\n", err)
	}
}

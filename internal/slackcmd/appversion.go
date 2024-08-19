package slackcmd

import (
	"github.com/A-K-2001/slack-test/database"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func AppCammonds(ev *slackevents.MessageEvent, client *slack.Client, repo *database.Repository) {

	// if strings.ToLower(strings.Split(ev.Text, " ")[1])  handle here all

	client.PostMessage(ev.Channel, slack.MsgOptionText("comming soon..", true))

	// bot.AddCommand(&slacker.CommandDefinition{
	// 	Command: "app -v -p {platform}",
	// 	Handler: func(ctx *slacker.CommandContext) {
	// 		var platform models.Platform = models.PlatformUNKNOWN
	// 		p := strings.ToUpper(ctx.Request().Param("platform"))
	// 		fmt.Println("get app version with platform", p)

	// 		switch p {
	// 		case "ANDROID":
	// 			platform = models.PlatformANDROID

	// 		case "IOS":
	// 			platform = models.PlatformIOS

	// 		case "SERVICE":
	// 			platform = models.PlatformSERVICE
	// 		}

	// 		res, err := repo.GetAppUpdatesByPlatform(ctx.Context(), platform)

	// 		if err != nil {
	// 			ctx.Response().PostError("error geting rider", err)
	// 		}
	// 		reply, err := json.MarshalIndent(res, "", "   ")
	// 		if err != nil {
	// 			ctx.Response().PostError("error geting rider data", err)
	// 		}
	// 		ctx.Response().Reply(string(reply))
	// 	},
	// 	HideHelp: false,
	// })

	// bot.AddCommand(&slacker.CommandDefinition{
	// 	Command: "app -v",
	// 	Handler: func(ctx *slacker.CommandContext) {
	// 		print("yes i am here ")
	// 		event := ctx.Event()
	// 		fmt.Println("test get appv", event)
	// 		res, err := repo.GetAppUpdates(ctx.Context())
	// 		if err != nil {
	// 			ctx.Response().PostError("error app-v", err)
	// 		}
	// 		reply, err := json.MarshalIndent(res, " ", "   ")
	// 		if err != nil {
	// 			ctx.Response().PostError("error geting app-v data", err)
	// 		}
	// 		ctx.Response().Reply(string(reply))
	// 	},
	// 	HideHelp: false,
	// })
}

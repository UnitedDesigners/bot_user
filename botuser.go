package main // import "github.com/uniteddesigners/botuser"

import (
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		// fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {

		case *slack.ConnectedEvent:
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello, I'm online now!", "G57SW7KEY"))

		case *slack.TeamJoinEvent:
			sendWelcome(api, ev.User.ID)

		default:
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}

package main // import "github.com/uniteddesigners/botuser"

import (
	"fmt"

	"github.com/nlopes/slack"
)

func sendWelcome(api *slack.Client, user string) {
	_, _, channelID, err := api.OpenIMChannel(user)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	welcomeMessageOptions := slack.NewPostMessageParameters()
	rulesOptions := slack.NewPostMessageParameters()
	professionalchannelOptions := slack.NewPostMessageParameters()
	informalChannelOptions := slack.NewPostMessageParameters()
	welcomeMessageOptions.LinkNames = 1
	rulesOptions.Attachments = []slack.Attachment{
		slack.Attachment{
			Fallback: "Be nice – racism, sexism, and any other forms of bigotry will result in immediate removal from the group.",
			Text:     "Be nice – racism, sexism, and any other forms of bigotry will result in immediate removal from the group.",
		},
		slack.Attachment{
			Fallback: "Don’t use United Designers as a platform to advertise (job postings may be posted in #jobs).",
			Text:     "Don’t use United Designers as a platform to advertise (job postings may be posted in <#C2D8Z0F46|jobs>).",
		},
		slack.Attachment{
			Fallback: "Please keep professional channels professional and on-topic. If a conversation becomes too off-topic, move it to an appropriate channel.",
			Text:     "Please keep professional channels professional and on-topic. If a conversation becomes too off-topic, move it to an appropriate channel.",
		},
	}
	professionalchannelOptions.Attachments = []slack.Attachment{
		slack.Attachment{
			Fallback: "#design: The general design channel",
			Title:    "<#C4MR6TD9T|design>",
			Text:     "The general design channel.",
		},
		slack.Attachment{
			Fallback: "#feedback-and-roasts: The general design channel",
			Title:    "<#C237ALERG|feedback-and-roasts>",
			Text:     "Post your work here to get feedback and critiques.",
		},
		slack.Attachment{
			Fallback: "Other professional channels: #frontend, #inspiration, #how-to, #prototyping, #software, #jobs",
			Text:     "Other professional channels: <#C37F1QTB2|frontend>, <#C1T80RU86|inspiration>, <#C5U3TQS76|accessibility>, <#C2A4UDC9Z|how-to>, <#C3RBYFVMG|prototyping>, <#C1YC4PS8L|software>, <#C2D8Z0F46|jobs>.",
		},
	}
	informalChannelOptions.Attachments = []slack.Attachment{
		slack.Attachment{
			Fallback: "#chit-chat: This is our general purpose chat channel (warning: occasionally semi-nsfw)",
			Title:    "<#C4Y3GAYK1|chit-chat>",
			Text:     "This is our general purpose chat channel (warning: occasionally semi-nsfw).",
		},
		slack.Attachment{
			Fallback: "Other channels: #safe-space, #vent, #films-and-tv, #nerd-cave, #music, #finance, #bayarea-events",
			Text:     "Other channels: <#C3VFC9HT6|safe-space>, <#C466HLB0X|vent>, <#C4W3TURD5|films-and-tv>, <#C26FVA9SN|nerd-cave>, <#C2H8CD141|music>, <#C5HJRL5U5|finance>, <#C23NJCEF7|bayarea-events>, and more.",
		},
	}

	api.PostMessage(channelID, "Welcome to United Designers! Go ahead and update your profile with a picture of yourself and say hi in #introductions!", welcomeMessageOptions)
	api.PostMessage(channelID, "To make sure that we remain an open and friendly environment, we have a few rules for our community here:", rulesOptions)
	api.PostMessage(channelID, "Here are a few of our professional channels:", professionalchannelOptions)
	api.PostMessage(channelID, "Here are a few of our other channels:", informalChannelOptions)
}

package usecase

import (
	"fmt"

	"github.com/rustysys-dev/attendance-tc/internal/utils/config"
	"github.com/slack-go/slack"
)

func SendMessage(msg string) {
	api := slack.New(config.SlackToken())
	channelID, timestamp, err := api.PostMessage(
		config.SlackChannelID(),
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		_, _ = fmt.Printf("%s\n", err)
		return
	}
	_, _ = fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}

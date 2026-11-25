package notifier

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

// SendMessage posts a message to Slack when the required environment variables are configured.
func SendMessage(message string) error {
	token := os.Getenv("DAROGA_SLACK_TOKEN")
	channelID := os.Getenv("DAROGA_SLACK_CHANNEL")
	if token == "" || channelID == "" {
		return fmt.Errorf("Slack notifications require DAROGA_SLACK_TOKEN and DAROGA_SLACK_CHANNEL")
	}

	api := slack.New(token)
	_, _, err := api.PostMessage(channelID, slack.MsgOptionText(message, false))
	if err != nil {
		return fmt.Errorf("send Slack notification: %w", err)
	}
	return nil
}

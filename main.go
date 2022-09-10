package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("xoxb-4075739431089-4086825624704-dYrjsHsEiHjDMPMFP7tgVsiu") //creating a connection with slack api
	channelID, timestamp, err := api.PostMessage(
		"C041S4CSF3P",
		slack.MsgOptionText("Hello World", false))

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent successfully %s at %s", channelID, timestamp)
}

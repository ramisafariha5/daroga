package internal

import (
	"fmt"

	"github.com/slack-go/slack"
)

func SlackNotify() {
	api := slack.New("xapp-1-A041XGKMRU4-4110589420017-9f4768309dcbab8d6fa2b125c76bfe556664eac15726e7ef497a6296535e85e3") //creating a connection with slack api
	checkauth, _ := api.AuthTest()
	channelID, timestamp, err := api.PostMessage("C041S4CSF3P", slack.MsgOptionText("Hello World", false))
	fmt.Printf("%+v", checkauth)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent successfully %s at %s", channelID, timestamp)
}

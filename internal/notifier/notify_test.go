package notifier

import "testing"

func TestSendMessageRequiresEnvironmentConfiguration(t *testing.T) {
	t.Setenv("DAROGA_SLACK_TOKEN", "")
	t.Setenv("DAROGA_SLACK_CHANNEL", "")

	if err := SendMessage("hello"); err == nil {
		t.Fatal("expected missing Slack configuration error")
	}
}

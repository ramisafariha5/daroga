package cmd

import (
	"strings"
	"testing"

	"github.com/ramisafariha5/daroga/internal/inspector"
)

func TestFormatSlackMessage(t *testing.T) {
	message := formatSlackMessage([]inspector.ContainerSummary{{
		Name: "api", ID: "1234567890ab", State: "running", Status: "Up 10 minutes",
	}})

	if !strings.Contains(message, "api") || !strings.Contains(message, "running") {
		t.Fatalf("unexpected message: %q", message)
	}
}

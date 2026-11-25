package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ramisafariha5/daroga/config"
	"github.com/ramisafariha5/daroga/internal/inspector"
	"github.com/ramisafariha5/daroga/internal/notifier"
	"github.com/spf13/cobra"
)

var darogaRun = &cobra.Command{
	Use:   "run",
	Short: "Inspect Docker containers",
	RunE: func(cmd *cobra.Command, args []string) error {
		containerConfig := config.GetContainerInfo()
		if containerConfig.Format != "table" && containerConfig.Format != "json" {
			return fmt.Errorf("unsupported output format %q: use table or json", containerConfig.Format)
		}

		client, err := inspector.InitDockerClient()
		if err != nil {
			return err
		}
		defer client.Close()

		selector := containerConfig.ContainerId
		if selector == "" {
			selector = containerConfig.ContainerName
		}

		ctx, cancel := context.WithTimeout(cmd.Context(), 10*time.Second)
		defer cancel()
		dockerInspector := inspector.NewDockerInspector(client)
		summaries, err := dockerInspector.Inspect(ctx, selector)
		if err != nil {
			return err
		}

		if containerConfig.Format == "json" {
			encoder := json.NewEncoder(cmd.OutOrStdout())
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(summaries); err != nil {
				return fmt.Errorf("encode JSON output: %w", err)
			}
		} else {
			printTable(cmd, summaries)
		}

		if containerConfig.Notify {
			message := formatSlackMessage(summaries)
			if err := notifier.SendMessage(message); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(darogaRun)
	containerInfo := config.GetContainerInfo()
	darogaRun.Flags().StringVarP(&containerInfo.ContainerName, "name", "n", "", "container name to inspect")
	darogaRun.Flags().StringVarP(&containerInfo.ContainerId, "id", "i", "", "container ID or prefix to inspect")
	darogaRun.Flags().StringVarP(&containerInfo.Format, "format", "f", "table", "output format: table or json")
	darogaRun.Flags().BoolVar(&containerInfo.Notify, "notify", false, "send inspected container status to Slack")
}

func printTable(cmd *cobra.Command, summaries []inspector.ContainerSummary) {
	out := cmd.OutOrStdout()
	fmt.Fprintf(out, "%-14s %-24s %-35s %-12s %s\n", "ID", "NAME", "IMAGE", "STATE", "STATUS")
	for _, summary := range summaries {
		fmt.Fprintf(out, "%-14s %-24s %-35s %-12s %s\n", summary.ID, summary.Name, summary.Image, summary.State, summary.Status)
	}
}

func formatSlackMessage(summaries []inspector.ContainerSummary) string {
	lines := make([]string, 0, len(summaries))
	for _, summary := range summaries {
		lines = append(lines, fmt.Sprintf("%s (%s): %s — %s", summary.Name, summary.ID, summary.State, summary.Status))
	}
	return "Docker inspection: " + strings.Join(lines, "\n")
}

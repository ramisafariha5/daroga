package cmd

import (
	"fmt"

	"github.com/ramisafariha5/daroga/config"
	"github.com/ramisafariha5/daroga/internal/inspector"
	"github.com/spf13/cobra"
)

var darogaRun = &cobra.Command{
	Use:   "run",
	Short: "run",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.GetContainerInfo()
		fmt.Println(config)
		client := inspector.InitDockerClient()
		dockerInspector := inspector.NewDockerInspector(config, client)
		dockerInspector.Inspect()
	},
}

func init() {
	rootCmd.AddCommand(darogaRun)
	containerInfo := config.GetContainerInfo()
	darogaRun.Flags().StringVarP(&containerInfo.ContainerName, "cname", "n", "", "container name to inspect")
	darogaRun.Flags().StringVarP(&containerInfo.ContainerId, "id", "i", "", "container id to inspect")
}

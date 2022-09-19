package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daroga",
	Short: "daroga",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
	}
}

func TestInit() {
	fmt.Print("Testing Init")

}

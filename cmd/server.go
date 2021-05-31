package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

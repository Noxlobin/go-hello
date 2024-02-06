package cmd

// The version of the command. I feel like its required for everything.

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the version number out to the screen",
	Long:  "its pretty self explanatory.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-hello: 0.1")
	},
}

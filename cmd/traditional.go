package cmd

// "Traditional" version of the hello command. I'm adding this as a subcommand as opposed to a flag. It's basically the same thing but without any punctuation

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traditionalCmd)
}

var description = "uses a traditional version of the message"

var traditionalCmd = &cobra.Command{
	Use:   "traditional",
	Short: description,
	Long:  description,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello cats")
	},
}

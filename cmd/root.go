package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "go-hello",
	Short: "Prints a friendly, customizable greeting.",
	Long:  `Go Hello is a application used to print a greeting used by most developers onto the screen. It's insipered by the "hello" program from GNU, execpt recreated in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, cats!") // Thats it. Thats the whole program. If I wasn't creating a CLI, I would've just straight up written the same line here in the main file.
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

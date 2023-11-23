package cmd

import (
	"curly/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for all the dependencies",
	Long:  `Check for all the dependencies and show the result.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for all the dependencies...")
		if utils.IsCurlInstalled() {
			fmt.Println("Curl is installed.")
			fmt.Println("Everything is fine.")
		} else {
			fmt.Println("Curl is not installed.")
			os.Exit(1)
		}
	},
}

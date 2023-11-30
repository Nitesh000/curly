package cmd

import (
	"curly/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Create bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create a config file for the curly at the ~/.curly/curly.json ",
	Long:  `Create a config for the additional output of the result from an api.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Create {
			err := utils.CreateConfigureFile()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

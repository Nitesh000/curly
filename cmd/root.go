package cmd

import (
	"curly/types"
	"curly/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Version      bool
	Source       string
	CreateSource string
)

var rootCmd = &cobra.Command{
	Use:   "curly",
	Short: "Curly is a CLI tool for making http requests",
	Long: `A fast and flexible CLI tool for makeing http 
      requests with the help of curl.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Version {
			fmt.Println("Curly version: ", types.VERSION)
			os.Exit(0)
		}
		if len(args) < 1 {
			fmt.Println("Plase provide the operation name")
			fmt.Println("Usage: curly <operation> <filename>")
			os.Exit(1)
		}
	},
}

var runFile = &cobra.Command{
	Use:   "run",
	Short: "Run from the json file",
	Long: `It's gonna run from the json file and  show the content of the file.
          add the file name after the command and run the command on the same directory as the json file is.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide the file name")
			log.Fatal("Usage: ./cobra-tut run <filename>")
		}

		// NOTE: add the configuration options from the file
		conf := types.CurlyConfig{
			AppConfig: types.NewConfig(),
		}
		types.AppConfigString = conf.AppConfig.CreateConfigureString()

		Source = args[0]
		_, err := utils.ReadFile(Source)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var createFile = &cobra.Command{
	Use:   "create",
	Short: "Create a new json file",
	Long:  `It's gonna create a new json file and write the content of the file.`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the file name")
			fmt.Println("Usage: curly create <filename>")
			os.Exit(1)
		}
		ext := strings.Split(args[0], ".")
		if ext[len(ext)-1] != "json" {
			fmt.Println("Please provide the json file name")
			fmt.Println("Usage: ./cobra-tut create <filename.json>")
			os.Exit(1)
		}
		CreateSource = args[0]
		utils.CreateFile(CreateSource)
	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&Version, "version", "V", false, "print version")
	configCmd.Flags().BoolVarP(&Create, "create", "c", false, "create config file")
	rootCmd.AddCommand(runFile)
	rootCmd.AddCommand(createFile)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(checkCmd)

	// NOTE: if the root command returns some error then the whole program just go boom.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

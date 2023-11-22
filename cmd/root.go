package cmd

import (
	"curly/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Verbose      bool
	Source       string
	CreateSource string
)

var rootCmd = &cobra.Command{
	Use:   "curly",
	Short: "Curly is a CLI tool for making http requests",
	Long: `A fast and flexible CLI tool for makeing http 
      requests with the help of curl.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Plase provide the operation name")
			log.Fatal("Usage: ./curly <operation> <filename>")
			os.Exit(1)
		}
	},
}

var readFile = &cobra.Command{
	Use:   "read",
	Short: "Read from the json file",
	Long: `It's gonna read from the json file and  show the content of the file.
          add the file name after the command and run the command on the same directory as the json file is.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide the file name")
			log.Fatal("Usage: ./cobra-tut read <filename>")
		}
		Source = args[0]

		val, _ := utils.ReadFile(Source)
		fmt.Println(val)
		if Verbose {
			fmt.Println("Running the read comamnd verbosely...")
		} else {
			fmt.Println("Running the read comamnd...")
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
			fmt.Println("Usage: ./cobra-tut create <filename>")
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
	readFile.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(readFile)
	rootCmd.AddCommand(createFile)

	// NOTE: if the root command returns some error then the whole program just go boom.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

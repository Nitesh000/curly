package cmd

import (
	"cobra-tut/utils"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	Verbose bool
	Source  string
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

func Execute() {
	readFile.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(readFile)

	// NOTE: if the root command returns some error then the whole program just go boom.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

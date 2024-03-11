package cmd

import (
	"fmt"
	"os"

	"github.com/anshitmishra/golangjsonfile/app/cmd/database"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	availableSubcommands = []*cobra.Command{}
)
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo app to demonstrate cobra",
	Long:  `demo app to demonstrate cobra by addition`,
}

var optionCmd = &cobra.Command{
	Use:   "options",
	Short: "demo app to demonstrate cobra",
	Long:  `demo app to demonstrate cobra by addition`,
	Run: func(cmd *cobra.Command, args []string) {
		Options()
	},
}

func init() {

	// Add more flags as needed
	availableSubcommands = []*cobra.Command{
		database.DatabaseCmd,
		optionCmd,
	}

	rootCmd.AddCommand(availableSubcommands...)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Options() {

	prompt := promptui.Select{
		Label: "Select an option",
		Items: []string{"database", "APIs", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "database":
		database.DatabaseOption()
	default:
		fmt.Println("saasdas")
	}
}

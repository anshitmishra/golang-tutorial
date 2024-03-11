/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package database

import (
	"fmt"
	"log"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// database/databaseCmd represents the database/database command
var DatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		flags, _ := cmd.Flags().GetString("create")
		if flags != "" {
			Created(flags)
		} else {
			DatabaseOption()
		}
	},
}

// CreateDbCmd represents the CreateDb command
var CreateDbCmd = &cobra.Command{
	Use:   "CreateDb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateDb called")
	},
}

var CreateDbName string

func init() {
	DatabaseCmd.PersistentFlags().StringVarP(&CreateDbName, "create", "c", "", "AWS region (required)")
	DatabaseCmd.AddCommand(CreateDbCmd)
}

func DatabaseOption() {
	prompt := promptui.Select{
		Label: "Select an database Options",
		Items: []string{"Create", "List", "Delete", "Rename"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Create":
		CreateDb()
		fmt.Println("saasdas")
	default:
		fmt.Println("saasdas")
	}

}

func CreateDb() {
	validate := func(input string) error {
		_, err := regexp.MatchString("[a-z][0-9]", "peach")

		return err
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Spicy Level",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}
	Created(result)
}

func Created(data string) {
	fmt.Println(data)

}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ahmetzumber/rapid-note-cli/internal/user"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rapid-note-cli",
	Short: "Rapid note provides you taking a notes dynamically from terminal.",
	Long:  `A longer description that rapid note loves you.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Rapid World !!")
	},
}

var createUserCmd = &cobra.Command{
	Use: "create",
	Short: "This command creates a new user.",
	Run: func(cmd *cobra.Command, args []string) {
		newUser := user.Create(args[0])
		fmt.Println("Welcome "+ newUser + " !")
	},

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rapid-note-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createUserCmd)
}

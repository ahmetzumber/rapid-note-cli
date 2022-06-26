package cmd

import (
	"fmt"
	"github.com/ahmetzumber/rapid-note-cli/internal/modal"
	"github.com/ahmetzumber/rapid-note-cli/internal/postgre"
	"github.com/ahmetzumber/rapid-note-cli/internal/repository"
	"os"

	"github.com/spf13/cobra"
)

type Launcher struct {
	Repo        		repository.IRepository
	CurrentUser			modal.User
	CurrentUserNote  	modal.Note
}

var LauncherObj Launcher

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
	Short: "This command creates a new modal.",
	Run: func(cmd *cobra.Command, args []string) {
		LauncherObj.Repo.AddUser(modal.User{
			Username: args[0],
			Email:    args[1],
		})
		LauncherObj.CurrentUser = modal.User{
			Username: args[0],
			Email:    args[1],
		}
		fmt.Println("Welcome "+ args[0] + " !")
	},
}

var userListCmd = &cobra.Command{
	Use: "list-users",
	Short: "This command list all users.",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := LauncherObj.Repo.GetUsers()
		if err != nil {
			fmt.Printf("User list error: %s", err)
		}
		for i := 0; i < len(users); i++ {
			fmt.Println("-> "+ users[i].Username)
		}
	},
}

var writeNoteCmd = &cobra.Command{
	Use: "write",
	Short: "With this command you can write your notes properly.",
	Run: func(cmd *cobra.Command, args []string) {
		LauncherObj.Repo.AddNote(modal.Note{
			UserID: LauncherObj.CurrentUser.ID,
			Data: args[0],
		})
		newNote := modal.Note{
			UserID: LauncherObj.CurrentUser.ID,
			Data: args[0],
		}
		LauncherObj.CurrentUserNote = newNote
	},
}

var test = &cobra.Command{
	Use: "test",
	Short: "test command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(LauncherObj.CurrentUser.ID)
		fmt.Println(LauncherObj.CurrentUser.Username)
		fmt.Println(LauncherObj.CurrentUser.Email)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	db, err := postgre.NewPostgreDB(postgre.Config)
	if err != nil {
		fmt.Printf("DB error: %s",err)
	}
	LauncherObj.Repo = repository.NewRepository(db)
	err = rootCmd.Execute()
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
	rootCmd.AddCommand(userListCmd)
	rootCmd.AddCommand(writeNoteCmd)
	rootCmd.AddCommand(test)
}

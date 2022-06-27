package cmd

import (
	"fmt"
	"github.com/ahmetzumber/rapid-note-cli/internal/modal"
	"github.com/ahmetzumber/rapid-note-cli/internal/postgre"
	"github.com/ahmetzumber/rapid-note-cli/internal/repository"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Launcher struct {
	Repo        		repository.IRepository
}

var LauncherObj Launcher

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
		WriteCurrentUser(args[0])
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
			UserID: LauncherObj.Repo.GetUserIDByUsername(GetCurrentUser()),	// you will write your note to current user
			Data: args[0],
		})
	},
}

var getCurrentUserNotesCmd = &cobra.Command{
	Use: "list-notes",
	Short: "This command list your all notes.",
	Run: func(cmd *cobra.Command, args []string) {
		currentUserID := LauncherObj.Repo.GetUserIDByUsername(GetCurrentUser())
		notes := LauncherObj.Repo.GetCurrentUserNotesByUserID(currentUserID)
		for i := 0; i < len(notes); i++ {
			fmt.Println("-> "+ notes[i].Data)
		}
	},
}

var test = &cobra.Command{
	Use: "test",
	Short: "test command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetCurrentUser())
		fmt.Println(LauncherObj.Repo.GetUserIDByUsername(GetCurrentUser()))
	},
}

func Execute() {
	db, err := postgre.NewPostgreDB(postgre.Config)
	if err != nil {
		fmt.Printf("DB error: %s",err)
	}
	LauncherObj.Repo = repository.NewRepository(db)
	repository.Migrate(db)
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func WriteCurrentUser(user string) error {
	d1 := []byte(user)
	f, err := os.Create("cmd/currentUser")
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile("cmd/currentUser", d1, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	return err
}

func GetCurrentUser() string {
	user, err := os.ReadFile("cmd/currentUser")
	if err != nil {
		log.Fatalln(err)
	}
	return string(user)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createUserCmd)
	rootCmd.AddCommand(userListCmd)
	rootCmd.AddCommand(writeNoteCmd)
	rootCmd.AddCommand(test)
	rootCmd.AddCommand(getCurrentUserNotesCmd)
}

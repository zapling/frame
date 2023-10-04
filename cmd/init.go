package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		if _, err := os.Stat(projectName); !os.IsNotExist(err) {
			fmt.Println("Directory already exists!")
			os.Exit(1)
		}

		if err := os.Mkdir(projectName, 0755); err != nil {
			fmt.Printf("Failed to create project directory: %v", err)
			os.Exit(1)
		}

		if err := os.Chdir(projectName); err != nil {
			fmt.Printf("Failed to change working directory to %s", projectName)
			os.Exit(1)
		}
	},
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zapling/gx/initialize"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Args:  cobra.ExactArgs(1),
	Run:   initialize.ExecuteCommand,
}

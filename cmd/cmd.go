package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zapling/frame/cmd/generate"
	"github.com/zapling/frame/cmd/initialize"
	"github.com/zapling/frame/cmd/serve"
	"github.com/zapling/frame/cmd/version"
)

var root = &cobra.Command{
	Use:   "frame",
	Short: "frame is a CLI tool to help developers write frontend applications in go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: You need to specify a command. Use '--help' to view the available commands.")
		os.Exit(1)
	},
}

func GetCommand() *cobra.Command {
	root.AddCommand(version.Command)
	root.AddCommand(initialize.Command)
	root.AddCommand(serve.Command)

	root.AddCommand(generate.Command)
	generate.Command.Flags().BoolVar(
		&generate.ComponentHandler,
		"with-handler",
		false,
		"Generate http handler",
	)

	return root
}

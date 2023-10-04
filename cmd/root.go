package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gx",
	Short: "gx is tool to help write frontend applications in go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from rootCmd")
	},
}

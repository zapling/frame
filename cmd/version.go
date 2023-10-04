package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show gx and dependecy versions",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("gx built without module, this should not be possible")
			os.Exit(1)
		}

		gxVersion := info.Main.Version
		if gxVersion == "" {
			gxVersion = "devel"
		}

		fmt.Printf("gx version: %s\n", gxVersion)
	},
}

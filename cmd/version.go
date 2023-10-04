package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version of gx",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("gx built without module, this should not be possible")
			os.Exit(1)
		}

		gxVersion := info.Main.Version
		if gxVersion == "" {
			gxVersion = "(devel)"
		}

		goVersion, _ := getInstalledGoVersion()
		if goVersion == "" {
			goVersion = "?"
		}

		fmt.Printf("gx version: %s\n", gxVersion)
		fmt.Printf("go version: %s\n", goVersion)
	},
}

func getInstalledGoVersion() (string, error) {
	goVersionCmdStdout := bytes.NewBuffer(nil)
	goVersionCmd := exec.Command("go", "version")
	goVersionCmd.Stdout = goVersionCmdStdout
	if err := goVersionCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run 'go version': %w", err)
	}

	goVersionBytes, _ := io.ReadAll(goVersionCmdStdout)
	goVersion := strings.Split(string(goVersionBytes), " ")[2][2:]

	return goVersion, nil
}

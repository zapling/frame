package version

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var errCommandNotFound = errors.New("command not found in path")

var Command = &cobra.Command{
	Use:   "version",
	Short: "Outputs frame CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("frame built without module, this should not be possible")
			os.Exit(1)
		}

		goVersion := getInstalledGoVersion()
		if goVersion == "" {
			goVersion = "?"
		}

		frameVersion := info.Main.Version
		if frameVersion == "" {
			frameVersion = "(devel)"
		}

		templVersion := getInstalledTemplVersion()
		if templVersion == "" {
			templVersion = "?"
		}

		fmt.Printf("go version: %s\n", goVersion)
		fmt.Printf("frame version: %s\n", frameVersion)
		fmt.Printf("templ version: %s\n", templVersion)
	},
}

func getInstalledGoVersion() string {
	output, err := getVersion([]string{"go", "version"})
	if errors.Is(err, errCommandNotFound) {
		return fmt.Sprint(chalk.Red, "'go' not found in path", chalk.Reset)
	}

	goVersion := strings.Split(output, " ")[2][2:]
	return goVersion
}

func getInstalledTemplVersion() string {
	output, err := getVersion([]string{"templ", "version"})
	if errors.Is(err, errCommandNotFound) {
		return fmt.Sprint(chalk.Red, "templ not found in path", chalk.Reset)
	}
	if err != nil {
		fmt.Printf("Failed to get templ version: %v", err)
		os.Exit(1)
	}

	return strings.TrimSpace(output)
}

func getVersion(command []string) (string, error) {
	if len(command) == 0 {
		return "", fmt.Errorf("empty command not allowed")
	}

	_, err := exec.LookPath(command[0])
	if err != nil {
		return "", errCommandNotFound
	}

	stdout := bytes.NewBuffer(nil)
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}

	bytes, err := io.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("failed to read bytes from stdout: %w", err)
	}

	return string(bytes), nil
}

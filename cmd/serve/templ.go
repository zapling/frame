package serve

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/ttacon/chalk"
)

func isUsingTempl() bool {
	cmd := exec.Command("go", "list", "-m", "-f", "{{ .Version }}", "github.com/a-h/templ")

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func generateTemplFiles() error {
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)

	cmd := exec.Command("templ", "generate")
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	fmt.Print("Generating templ files... ")

	if err := cmd.Run(); err != nil {
		fmt.Print(chalk.Red, "✘\n", chalk.Reset)

		bytes, _ := io.ReadAll(stdout)
		fmt.Print(chalk.Red, string(bytes), err, "\n", chalk.Reset)

		return fmt.Errorf("failed to generate templ files: %w", err)
	}

	fmt.Print(chalk.Green, "✔\n", chalk.Reset)

	return nil
}

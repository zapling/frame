package serve

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/ttacon/chalk"
)

func getAppCmd(binaryPath string) *exec.Cmd {
	stdout := bytes.NewBuffer(nil)
	cmd := exec.Command(binaryPath)
	cmd.Stdout = stdout
	return cmd
}

func runApp(cmd *exec.Cmd) error {
	fmt.Print("Starting app... ")

	if err := cmd.Start(); err != nil {
		fmt.Print(chalk.Red, "✘\n", chalk.Reset)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic("this is weird: " + err.Error())
		}

		bytes, _ := io.ReadAll(stdout)
		fmt.Print(chalk.Red, string(bytes), err, "\n", chalk.Reset)

		return err
	}

	fmt.Print(chalk.Green, "✔\n", chalk.Reset)

	return nil
}

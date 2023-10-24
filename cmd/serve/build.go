package serve

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os/exec"

	"github.com/ttacon/chalk"
)

func buildApp() (binaryPath string, err error) {
	if isUsingTempl() {
		err := generateTemplFiles()
		if err != nil {
			return "", fmt.Errorf("failed to generate templ files: %w", err)
		}
	}

	binaryPath, err = buildAppBinary()
	if err != nil {
		return "", fmt.Errorf("failed to build app: %w", err)
	}

	return
}

func buildAppBinary() (binaryPath string, err error) {
	path := "/tmp/frame-binary-" + getSemiRandomString()

	buf := bytes.NewBuffer(nil)
	buildCmd := exec.Command("go", "build", "-o", path, "./cmd/")
	buildCmd.Stderr = buf

	fmt.Print("Building app... ")

	if err := buildCmd.Run(); err != nil {
		fmt.Print(chalk.Red, "✘\n", chalk.Reset)

		bytes, _ := io.ReadAll(buf)
		fmt.Print(chalk.Red, string(bytes), err, "\n", chalk.Reset)

		return "", err
	}

	fmt.Print(chalk.Green, "✔\n", chalk.Reset)

	return path, nil
}

func getSemiRandomString() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

package serve

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/rjeczalik/notify"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var Command = &cobra.Command{
	Use:   "serve",
	Short: "Builds and serves your application, rebuilding on file changes",
	Run:   executeCommand2,
}

func executeCommand2(cmd *cobra.Command, args []string) {
	var (
		serveCmd *exec.Cmd
		err      error
	)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	watchFileChanged := make(chan notify.EventInfo, 1)
	if err := notify.Watch("./...", watchFileChanged, notify.All); err != nil {
		fmt.Printf("Error: Failed to watch directory for changes: %v\n", err)
		os.Exit(1)
	}
	defer notify.Stop(watchFileChanged)

	binaryPath, err := buildApp()
	if err != nil {
		os.Exit(1)
	}

	serveCmd, err = runApp(binaryPath)
	if err != nil {
		os.Exit(1)
	}

	eventProcessedAt := time.Now()

	for {
		select {
		case <-watchFileChanged:

			// we might get multiple events regarding the same thing in bursts,
			// ignore event if we just rebuilt
			if !time.Now().Add(-1 * time.Second).After(eventProcessedAt) {
				break
			}

			binaryPath, err := buildApp()
			if err != nil {
				eventProcessedAt = time.Now()
				break
			}

			if serveCmd != nil {
				if err := serveCmd.Process.Signal(syscall.SIGTERM); err != nil {
					fmt.Printf("Error stopping app: %v\n", err)
					break
				}

				if err := serveCmd.Wait(); err != nil {
					fmt.Printf("Error waiting for app: %v\n", err)
					break
				}
			}

			serveCmd, err = runApp(binaryPath)
			if err != nil {
				eventProcessedAt = time.Now()
				break
			}

			eventProcessedAt = time.Now()

		case <-osSignals:
			fmt.Println("Goodbye!")
			// TODO: santify check and try and kill all pids that we know we have created in the past?
			os.Exit(0)
		}
	}
}

func runApp(binaryPath string) (*exec.Cmd, error) {

	buf := bytes.NewBuffer(nil)
	command := exec.Command(binaryPath)
	command.Stdout = os.Stdout
	command.Stderr = buf

	fmt.Print("Starting app... ")

	if err := command.Start(); err != nil {
		fmt.Print(chalk.Red, "✘\n", chalk.Reset)

		bytes, _ := io.ReadAll(buf)
		fmt.Print(chalk.Red, string(bytes), err, "\n", chalk.Reset)

		return nil, fmt.Errorf("failed to start app: %w", err)
	}

	fmt.Print(chalk.Green, "✔\n", chalk.Reset)

	return command, nil
}

func buildApp() (string, error) {
	path := "/tmp/frame-binary-" + getSemiRandomString()

	buf := bytes.NewBuffer(nil)
	buildCmd := exec.Command("go", "build", "-o", path, "./cmd/")
	buildCmd.Stderr = buf

	fmt.Print("Building app... ")

	if err := buildCmd.Run(); err != nil {
		fmt.Print(chalk.Red, "✘\n", chalk.Reset)

		bytes, _ := io.ReadAll(buf)
		fmt.Print(chalk.Red, string(bytes), err, "\n", chalk.Reset)

		return "", fmt.Errorf("failed to build app: %w", err)
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

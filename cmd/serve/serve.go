package serve

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/rjeczalik/notify"
	"github.com/spf13/cobra"
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
		fmt.Printf("Error building app: %v\n", err)
		os.Exit(1)
	}

	serveCmd, err = runApp(binaryPath)
	if err != nil {
		fmt.Printf("Error running app: %v\n", err)
		os.Exit(1)
	}

	for {
		select {
		case <-watchFileChanged:
			binaryPath, err := buildApp()
			if err != nil {
				fmt.Printf("Error building app: %v\n", err)
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
				fmt.Printf("Error starting app: %v\n", err)
				break
			}

		case <-osSignals:
			fmt.Println("Goodbye!")
			// TODO: santify check and try and kill all pids that we know we have created in the past?
			os.Exit(0)
		}
	}
}

func runApp(binaryPath string) (*exec.Cmd, error) {

	command := exec.Command(binaryPath)
	command.Stdout = os.Stdout

	if err := command.Start(); err != nil {
		return nil, fmt.Errorf("failed to start app: %w", err)
	}

	return command, nil
}

func buildApp() (string, error) {
	path := "/tmp/frame-binary-" + getSemiRandomString()
	buildCmd := exec.Command("go", "build", "-o", path, "./cmd/")
	if err := buildCmd.Run(); err != nil {
		return "", err
	}

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

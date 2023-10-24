package serve

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/rjeczalik/notify"
	"github.com/spf13/cobra"
	"github.com/zapling/frame/pkg/cfg"
	"golang.org/x/net/websocket"
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

	os.Setenv(cfg.FrameDevMode, "true")

	devServer := &developmentServer{
		wsConns: make(map[*websocket.Conn]bool),
	}

	go devServer.start()

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

	serveCmd = getAppCmd(binaryPath)
	if err := runApp(serveCmd); err != nil {
		os.Exit(1)
	}

	devServer.notifyClients()

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

			serveCmd = getAppCmd(binaryPath)
			if err := runApp(serveCmd); err != nil {
				eventProcessedAt = time.Now()
				break
			}

			devServer.notifyClients()

			eventProcessedAt = time.Now()

		case <-osSignals:
			fmt.Println("Goodbye!")
			// TODO: santify check and try and kill all pids that we know we have created in the past?
			os.Exit(0)
		}
	}
}

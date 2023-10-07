package run

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Webserver(server *http.Server) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
			return
		}
	}()

	stopSignalChan := make(chan os.Signal, 1)
	defer close(stopSignalChan)
	signal.Notify(stopSignalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stopSignalChan:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown gracefully: %w", err)
		}

		return nil
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("failure when trying to serve: %w", err)
		}
		return nil
	}
}

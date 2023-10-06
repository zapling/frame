package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := startRouter(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-osSignals
	close(osSignals)
}

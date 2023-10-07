package main

import (
	"fmt"
	"os"

	"github.com/zapling/frame/pkg/run"
)

func main() {
	if err := run.Webserver(getServer()); err != nil {
		fmt.Printf("Webserver stopped: %v\n", err)
		os.Exit(1)
	}
}

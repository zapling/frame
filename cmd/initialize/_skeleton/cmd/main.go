package main

import (
	"fmt"
	"os"

	"github.com/zapling/frame/pkg/cfg"
	"github.com/zapling/frame/pkg/run"
)

func main() {
	frameConfig, err := cfg.Get()
	if err != nil {
		fmt.Printf("Failed to load frame config: %v\n", err)
		os.Exit(1)
	}

	if err := run.Webserver(getServer(frameConfig)); err != nil {
		fmt.Printf("Webserver stopped: %v\n", err)
		os.Exit(1)
	}
}

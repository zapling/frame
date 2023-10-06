package main

import (
	"fmt"
	"os"

	"github.com/zapling/gx/cmd"
)

func main() {
	if err := cmd.GetCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

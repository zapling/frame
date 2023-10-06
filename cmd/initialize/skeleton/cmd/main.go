package main

import (
	"fmt"
	"os"
)

func main() {
	if err := startRouter(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

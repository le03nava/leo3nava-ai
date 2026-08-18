package main

import (
	"os"

	"github.com/leo3nava/leo3nava_ai/tools-go/leo3nava-ai/internal/app"
)

func main() {
	err := app.Run(os.Stdout, os.Stderr, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}

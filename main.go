package main

import (
	"os"

	"gitlab.com/dl7850949/blob-storage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}

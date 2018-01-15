package main

import (
	"os"

	"github.com/gsantandrea/catbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

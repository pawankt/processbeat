package main

import (
	"os"

	"github.com/pawankt/processbeat/cmd"

	_ "github.com/pawankt/processbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

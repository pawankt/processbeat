package main

import (
	"os"

	"github.com/pk-devops/processbeat/cmd"

	_ "github.com/pk-devops/processbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"github.com/darianJmy/fortis-services/cmd/app"
	"os"
)

func main() {
	cmd := app.NewFortisServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

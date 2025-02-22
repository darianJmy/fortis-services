package main

import (
	"os"

	"github.com/darianJmy/fortis-services/cmd/app"
)

func main() {
	cmd := app.NewFortisServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

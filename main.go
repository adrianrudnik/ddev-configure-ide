package main

import (
	"github.com/adrianrudnik/ddev-configure-ide/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Configure a human-readable CLI output as this will run on DDEV lifecycle CLI commands
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stderr,
		PartsExclude: []string{"time"},
	})

	cmd.Execute()
}

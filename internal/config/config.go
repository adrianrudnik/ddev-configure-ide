package config

import (
	"github.com/adrianrudnik/ddev-configure-ide/internal/ddev"
	"github.com/rs/zerolog/log"
	"os"
)

func newConfig() *RuntimeConfig {
	return &RuntimeConfig{
		DryRun:           false,
		WorkingDirectory: "",
	}
}

func NewDefaultRuntimeConfig() *RuntimeConfig {
	return newConfig()
}

func (rc *RuntimeConfig) BootWorkingDirectory(suggestedWorkingDirectory string) {
	if suggestedWorkingDirectory != "" {
		// Check the given suggestion for validity
		fileInfo, err := os.Stat(suggestedWorkingDirectory)

		if err != nil {
			log.Warn().Err(err).Str("root-path", suggestedWorkingDirectory).Msg("Could not resolve suggested working directory, falling back to dot.")
			rc.WorkingDirectory = "."
		} else if !fileInfo.IsDir() {
			log.Warn().Err(err).Str("root-path", suggestedWorkingDirectory).Msg("Given root-path could not be identified as directory, falling back to dot.")
			rc.WorkingDirectory = "."
		}

		// Seems fine, lets take it and fail later
		rc.WorkingDirectory = suggestedWorkingDirectory
	} else {
		// Try to detect it through the OS and fall back to dot in the worst case
		if rc.WorkingDirectory == "" {
			v, err := os.Getwd()
			if err != nil {
				log.Warn().Err(err).Msg("Could not determine working directory through OS, falling back to dot.")
				rc.WorkingDirectory = "."
			} else {
				rc.WorkingDirectory = v
			}
		}
	}

	log.Info().Str("cwd", rc.WorkingDirectory).Msg("Resolved current working directory")

	// We need to be able to describe the current project that resides within the current working directory
	// or there is nothing we can do.
	rc.DDEVConfig = ddev.MustDescribeConfig(rc.WorkingDirectory)
}

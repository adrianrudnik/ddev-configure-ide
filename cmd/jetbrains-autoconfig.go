package cmd

import (
	"github.com/adrianrudnik/ddev-configure-ide/internal/config"
	"github.com/adrianrudnik/ddev-configure-ide/internal/datasources"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var jetbrainsAutoconfigCmd = &cobra.Command{
	Use:   "autoconfig",
	Short: "Tries to refresh the IDE configuration automatically",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := cmd.Flags().GetString("root-path")
		if err != nil {
			log.Error().Err(err).Msg("Unable to parse root-path flag")
			os.Exit(1)
		}

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Error().Err(err).Msg("Unable to parse dry-run flag")
			os.Exit(1)
		}

		log.Info().Msg("Configuring IDE")

		conf := config.NewDefaultRuntimeConfig()
		conf.BootWorkingDirectory(cwd)
		conf.DryRun = dryRun

		datasources.MustHave(conf)
	},
}

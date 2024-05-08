package cmd

import "github.com/spf13/cobra"

var jetbrainsCmd = &cobra.Command{
	Use:   "jetbrains",
	Short: "IDE helpers for JetBrains products",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

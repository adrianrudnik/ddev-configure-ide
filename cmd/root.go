package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	jetbrainsAutoconfigCmd.ResetFlags()
	jetbrainsAutoconfigCmd.ResetCommands()
	jetbrainsAutoconfigCmd.Flags().String("root-path", "", "Root path that contains the .idea folder")
	jetbrainsAutoconfigCmd.Flags().Bool("dry-run", false, "Dry run, print, do not actually write to configs")

	jetbrainsCmd.AddCommand(jetbrainsAutoconfigCmd)

	rootCmd.AddCommand(jetbrainsCmd)

}

func Execute() {
	rootCmd.SetOut(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "ddev-configure-ide",
	Short: "Show usage information",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

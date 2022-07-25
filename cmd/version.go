package cmd

import (
	flunky "github.com/francoarendholz/flunky/base"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of flunky",
	Run: func(cmd *cobra.Command, args []string) {
		flunky.PrintVersion()
	},
}

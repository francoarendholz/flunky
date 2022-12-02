package cmd

import (
	"github.com/francoarendholz/flunky/manage"
	"github.com/spf13/cobra"
)

var forceApprovePendingSignatures bool

func init() {
	rootCmd.AddCommand(manageCmd)
	manageCmd.AddCommand(systemMessageCmd)
	manageCmd.AddCommand(approvePendingSignaturesCmd)
	approvePendingSignaturesCmd.Flags().BoolVarP(&forceApprovePendingSignatures, "force", "f", false, "Force approval of ALL signatures (use with caution!)")
	manageCmd.AddCommand(decodeAllSecretsCmd)
}

var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Jenkins management and configuration",
	Long:  `Some tools for day to day Jenkins management and configuration tasks.`,
}

var systemMessageCmd = &cobra.Command{
	Use:   "systemMessage",
	Short: "Set Jenkins System Message.",
	Long:  `Set the System Welcome Message of the Jenkins Server.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manage.SetSystemMessage(args[0])
	},
}

var approvePendingSignaturesCmd = &cobra.Command{
	Use:   "approvePendingSignatures",
	Short: "Approve pending signatures.",
	Long:  `Approve all pending script signatures waiting for approval in Jenkins.`,
	Run: func(cmd *cobra.Command, args []string) {
		manage.ApprovePendingSignatures(forceApprovePendingSignatures)
	},
}

var decodeAllSecretsCmd = &cobra.Command{
	Use:   "decodeAllSecrets",
	Short: "Decode all secrets.",
	Long:  `Decode all secrets of Jenkins including folder and org level.`,
	Run: func(cmd *cobra.Command, args []string) {
		manage.DecodeAllSecrets()
	},
}

package cmd

import (
	"github.com/francoarendholz/flunky/jobs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(jobsCmd)
	jobsCmd.AddCommand(runPipelineScript)
	jobsCmd.AddCommand(getJobConfigXml)
}

var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Running and manipulating Jenkins jobs",
	Long:  `Some tools for running an manipulating jobs on an existing Jenkins instance.`,
}

var runPipelineScript = &cobra.Command{
	Use:   "runPipelineScript",
	Short: "Run a pipeline script.",
	Long:  `Execute a local pipeline script on a remote Jenkins instance. Provide a filepath!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobs.RunPipelineScript(args[0])
	},
}

var getJobConfigXml = &cobra.Command{
	Use:   "getJobConfigXml",
	Short: "Get the config.xml of a job.",
	Long:  `Get the config.xml of a job and print it to stdout.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobs.GetJobConfigXml(args[0])
	},
}

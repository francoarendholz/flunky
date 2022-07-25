/*
Package cmd for Flunky Jenkins toolkit
Copyright © 2022 Franco Arendholz <franco.arendholz@agile-rcm.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flunky",
	Short: "A command-line toolbox for Jenkins CI",
	Long:  `Flunky is a command-line toolbox for the Jenkins CI server.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	var jenkinsAPIUrl string
	var jenkinsAPIUser string
	var jenkinsAPIKey string
	var verbose bool

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.flunky.yaml)")

	rootCmd.PersistentFlags().StringVarP(&jenkinsAPIUrl, "jenkins_api_url", "s", "JENKINS_API_SERVER", "Jenkins server bese URL (e.g. https://ci.code-ops.eu/).")
	rootCmd.MarkFlagRequired("jenkins_api_url")
	viper.BindPFlag("jenkins_api_url", rootCmd.PersistentFlags().Lookup("jenkins_api_url"))

	rootCmd.PersistentFlags().StringVarP(&jenkinsAPIUser, "jenkins_api_user", "u", "JENKINS_API_USER", "Jenkins API enabled user to perform tasks.")
	rootCmd.MarkFlagRequired("jenkins_api_user")
	viper.BindPFlag("jenkins_api_user", rootCmd.PersistentFlags().Lookup("jenkins_api_user"))

	rootCmd.PersistentFlags().StringVarP(&jenkinsAPIKey, "jenkins_api_key", "k", "JENKINS_API_KEY", "Jenkins API key for API user.")
	rootCmd.MarkFlagRequired("jenkins_api_key")
	viper.BindPFlag("jenkins_api_key", rootCmd.PersistentFlags().Lookup("jenkins_api_key"))

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Output for all commands.")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".flunky" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".flunky")
	}

	viper.SetEnvPrefix("flunky")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

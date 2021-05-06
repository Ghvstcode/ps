package cmd

import (
	"github.com/spf13/cobra"
	"os"
)


var rootCmd = &cobra.Command{
	Use:           "ps",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Ps is a CLI tool for playing around with processes on your machine!",
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	rootCmd.AddCommand(getProcesses)
	rootCmd.AddCommand(getGoProcesses)
}
package cmd

import (
	"github.com/Ghvstcode/process/pkg/killGoProcess"
	"github.com/spf13/cobra"
)

var KillAllGoProcesses = &cobra.Command{
	Use:   "kill-go",
	Short: "Kill all the running GO Processes on your machine!",
	Long:  `Kill all the running GO Processes on your machine! This command returns a path to the executable file of these processes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := killGoProcess.KillGoProcess()
		return err
	},
}

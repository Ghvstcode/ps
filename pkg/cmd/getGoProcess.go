package cmd

import (
	"github.com/Ghvstcode/process/pkg/getprocess"
	"github.com/spf13/cobra"
)

var getGoProcesses = &cobra.Command{
	Use:   "list-go",
	Short: "Get all the running Processes on your machine!",
	Long: `Get all the running Processes on your machine! This command returns a path to the executable file of these processes.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		err := getprocess.ListGoProcess()
		return err
	},
}

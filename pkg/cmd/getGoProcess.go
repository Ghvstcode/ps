package cmd

import (
	"github.com/Ghvstcode/process/pkg/listGoprocess"
	"github.com/spf13/cobra"
)

var getGoProcesses = &cobra.Command{
	Use:   "list-go",
	Short: "Get all the running GO Processes on your machine!",
	Long: `Get all the running GO Processes on your machine! This command returns a path to the executable file of these processes.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		err := listGoprocess.ListGoProcess()
		return err
	},
}

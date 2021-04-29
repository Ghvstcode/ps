package cmd

import (
	"github.com/Ghvstcode/process/pkg/getprocess"
	"github.com/spf13/cobra"
)

var getProcesses = &cobra.Command{
	Use:   "list",
	Short: "Get all the running Processes on your machine!",
	Long: `Get all the running Processes on your machine! This command returns a path to the executable file of these processes.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		err := getprocess.ListAllProcesses()
		return err
	},
}


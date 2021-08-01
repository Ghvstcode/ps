package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "ps",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Ps is a CLI tool for playing around with processes on your machine!",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if strings.Contains(err.Error(), "unknown command") {
			s := rootCmd.SuggestionsFor(os.Args[1])
			if len(s) > 0 {
				fmt.Printf("Unknown Command, Did you mean \"%s\"?\nIf not, ", s[0])
			}
			fmt.Println("run the --help flag to get a list of all commands")
		}
		os.Exit(-1)
	}
}

func init() {
	rootCmd.AddCommand(getProcesses)
	rootCmd.AddCommand(getGoProcesses)
	rootCmd.AddCommand(KillAllGoProcesses)
}

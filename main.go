package main

import (
	"os"

	"github.com/Januadrym/rnote/command"

	"github.com/spf13/cobra"
)

func main() {
	listCmd := &cobra.Command{
		Use:   "Rnote",
		Short: "Welcome back, Quan",
	}

	listCmd.AddCommand(command.GetTime())
	listCmd.AddCommand(command.CreateNewNote())
	listCmd.SetHelpCommand(command.SetHelp())
	if err := listCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

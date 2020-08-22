package main

import (
	"os"

	"github.com/Januadrym/rnote/command"

	"github.com/spf13/cobra"
)

func main() {
	listCmd := &cobra.Command{
		Use:   "Rnote",
		Short: "Welcome back, " + command.GetUserName(),
	}

	listCmd.AddCommand(command.GetTime())
	listCmd.AddCommand(command.ListAllNote())
	listCmd.AddCommand(command.CreateNewNote())
	listCmd.AddCommand(command.ReadNote())
	listCmd.AddCommand(command.EditNote())
	listCmd.AddCommand(command.DeleteNote())
	listCmd.SetHelpCommand(command.SetHelp())

	if err := listCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	listCmd := &cobra.Command{
		Use:   "Rnote",
		Short: "Welcome back, Quan",
	}

	listCmd.AddCommand(getTimeCmd())
	if err := listCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

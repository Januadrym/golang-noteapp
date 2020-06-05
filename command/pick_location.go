package command

import "github.com/spf13/cobra"

func directory() *cobra.Command {
	return &cobra.Command{
		Use:   "dir",
		Short: "Pick directory for saving note, otherwise at default",
	}
}

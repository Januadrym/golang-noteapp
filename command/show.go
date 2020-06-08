package command

import (
	"github.com/spf13/cobra"
)

// ListAllNote list all the note you have in current directory
func ListAllNote() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show all note",
		Long:  "Display all note you have in current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := showListFile(); err != nil {
				return err
			}
			return nil
		},
	}
}

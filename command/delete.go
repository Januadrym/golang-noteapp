package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// DeleteNote delete one existing note
func DeleteNote() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete one note",
		Long:  "Choose one existing note to delete",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := showListFile(); err != nil {
				return err
			}
			fmt.Println("Choose one note to delete")
			filename, err := inputFileName()
			if err != nil {
				return err
			}
			if _, er := os.Stat(filename); os.IsNotExist(er) {
				fmt.Println("File does not exist")
				return er
			}

			// delete
			if err := os.Remove(filename); err != nil {
				return err
			}
			fmt.Println("Note deleted, name:", filename)
			return nil
		},
	}
}

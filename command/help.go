package command

import "github.com/spf13/cobra"

// SetHelp display help to use the CLI
func SetHelp() *cobra.Command {
	return &cobra.Command{
		Use: "howto",
		Long: `In case you don't know how to use.
I'm here to help.

Here is how:
	Chose the location to store your note
	"new" to create a new note, !exit! to save and quit
	"read" to open a existing note
	"edit" to edit - not really, can only append to new text, !exit! to save and quit
	"delete" to delete, obviously
	"show" to show all *.txt file in current directory
	"time" addition command to display time 

Furthur regards,
Quan`,
	}
}

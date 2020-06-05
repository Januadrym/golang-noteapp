package command

import "github.com/spf13/cobra"

// SetHelp display help to use the CLI
func SetHelp() *cobra.Command {
	return &cobra.Command{
		Use: "howto",
		Long: `In case you don't know how to use.
I'm here to help.
And you're pretty dumb :)

Furthur Regards,
	Quan`,
	}
}

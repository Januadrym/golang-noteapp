package command

import (
	"time"

	"github.com/spf13/cobra"
)

// GetTime command to tell time
func GetTime() *cobra.Command {
	return &cobra.Command{
		Use:   "timenow",
		Short: "Telling the time",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hello Quan, the time is: ", prettyTime)
			return nil
		},
	}
}

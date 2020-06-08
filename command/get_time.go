package command

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// GetTime command to tell time
func GetTime() *cobra.Command {
	return &cobra.Command{
		Use:   "time",
		Short: "Telling the time",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format("15:04:05, Monday, January 02 2006")
			fmt.Println("Hello Quan, the time is: ", prettyTime)
			return nil
		},
	}
}

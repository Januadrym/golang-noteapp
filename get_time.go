package main

import (
	"time"

	"github.com/spf13/cobra"
)

func getTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "timenow",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hello Quan, the time is: ", prettyTime)
			return nil
		},
	}
}

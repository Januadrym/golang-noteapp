package command

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// ReadNote choose an existing note to read
func ReadNote() *cobra.Command {
	return &cobra.Command{
		Use:   "read",
		Short: "Read an existing note",
		Long:  "Choose an existing note to read :|",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := showListFile(); err != nil {
				return nil
			}
			fmt.Println("Choose a note to read")
			filename, err := inputFileName()
			if err != nil {
				return err
			}
			if _, er := os.Stat(filename); os.IsNotExist(er) {
				fmt.Println("File does not exist")
				return er
			}

			file, er := os.Open(filename)
			if er != nil {
				fmt.Printf("Error opening %s: %s", filename, er)
				return er
			}
			defer file.Close()

			// read
			fmt.Println("Content in", filename, ">>")
			reader := bufio.NewReader(file)
			for {
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Printf("error reading file %s", err)
					break
				}
				fmt.Print(line)
			}
			return nil
		},
	}
}

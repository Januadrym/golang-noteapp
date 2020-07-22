package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// CreateNewNote create a new note with a file name
func CreateNewNote() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Create new note at current directory",
		Long:  "Create new note with file's name\nFile name must be longer than 3 character\nTo finished and save, type in: !exit! ",
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, err := inputFileName()
			if err != nil {
				fmt.Println("Oops, you did something wrong")
				return err
			}
			if _, er := os.Stat(filename); os.IsNotExist(er) {
				fmt.Println(filename, "created at current directory.")
				fmt.Println("Type in: !exit! to save and quit")
			} else {
				fmt.Println("File already exist")
				return er
			}
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			// mark save point
			defer fmt.Fprint(file, "-------> At ", time.Now().Format("Mon Jan 2 15:04:05"))

			// write
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("-> ")

				text, _ := reader.ReadString('\n')
				finaltext, err := trimEndLineSequence(text)
				if err != nil {
					return err
				}

				if strings.Compare(finaltext, "!exit!") == 0 {
					break
				}
				fmt.Fprintf(file, finaltext+"\n")
			}
			return nil
		},
	}
}

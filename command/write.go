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
		Short: "Create new note",
		Long:  "Create new note with file name\nTo finish type: !exit! ",
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, err := inputFileName()
			if err != nil {
				fmt.Print("Error: ", err)
			}
			fmt.Println(filename, "created at current directory.")

			file, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}

			defer file.Close()
			defer fmt.Fprint(file, "-------> At ", time.Now().Format("Mon Jan 2 15:04:05"))

			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("-> ")

				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\r\n", "", -1)

				if strings.Compare(text, "!exit!") == 0 {
					break
				}
				fmt.Fprintf(file, text+"\n")
			}
			return nil
		},
	}
}

func inputFileName() (string, error) {
	inputFileName := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name: ")
	filename, _ := inputFileName.ReadString('\n')
	filename = strings.Replace(filename, "\r\n", "", -1)

	if filename == "" {
		return "", fmt.Errorf("No file name")
	}
	return filename + ".txt", nil
}

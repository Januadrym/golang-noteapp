package command

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// EditNote delete one existing note
func EditNote() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit one note",
		Long:  "Choose one existing note to put in more text",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := showListFile(); err != nil {
				return err
			}
			fmt.Println("Choose one note to edit")
			filename, err := inputFileName()
			if err != nil {
				return err
			}
			if _, er := os.Stat(filename); os.IsNotExist(er) {
				fmt.Println("File does not exist")
				return er
			}

			// edit
			file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
			if err != nil {
				return err
			}
			defer file.Close()
			defer fmt.Fprint(file, "\n-------> At ", time.Now().Format("Mon Jan 2 15:04:05"))

			// old data
			fmt.Println("Content in:", filename)
			reader := bufio.NewReader(file)
			for {
				line, err := reader.ReadString('\n')
				fmt.Print(line)
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Printf("error reading file %s", err)
					break
				}
			}
			// appending new data
			inputR := bufio.NewReader(os.Stdin)
			fmt.Println("")
			for {
				fmt.Print("-> ")

				text, _ := inputR.ReadString('\n')
				text = strings.Replace(text, "\r\n", "", -1)

				if strings.Compare(text, "!exit!") == 0 {
					break
				}
				fmt.Fprintf(file, "\n"+text)
			}
			return nil
		},
	}
}

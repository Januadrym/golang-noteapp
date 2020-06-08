package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func showListFile() error {
	wd, _ := os.Getwd()
	var files []string
	err := filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(wd) == ".txt" {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(files) != 0 {
		fmt.Println("Here are all notes: ")
		fmt.Println("+------------------+")
		for _, file := range files {
			fmt.Println("    ", file)
		}
		fmt.Println("+------------------+")
	} else {
		fmt.Println("You don't have any note here")
	}
	return nil
}

func inputFileName() (string, error) {
	inputFileName := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name: ")
	filename, _ := inputFileName.ReadString('\n')
	filename = strings.Replace(filename, "\r\n", "", -1)

	if len(filename) < 1 {
		return "", errors.New("Please input correct file name")
	}
	return filename + ".txt", nil
}

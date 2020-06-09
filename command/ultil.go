package command

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func showListFile() error {
	wd, _ := os.Getwd()
	files, err := ioutil.ReadDir(wd)
	if err != nil {
		return err
	}

	if len(files) != 0 {
		fmt.Println("Here are all notes: ")
		fmt.Println("+------------------+")
		for _, file := range files {
			if check, _ := regexp.MatchString(".txt", file.Name()); check {
				fmt.Println("    ", file.Name())
			}
		}
		fmt.Println("+------------------+")
	} else {
		fmt.Println("You don't have any note here")
		return errors.New("Empty")
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

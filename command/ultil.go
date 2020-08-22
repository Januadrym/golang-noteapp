package command

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func showListFile() error {
	wd, _ := os.Getwd()
	files, err := ioutil.ReadDir(wd)
	if err != nil {
		return err
	}
	var txts []string
	for _, file := range files {
		if check, _ := regexp.MatchString(".txt", file.Name()); check {
			txts = append(txts, file.Name())
		}
	}

	if len(txts) == 0 {
		fmt.Println("You don't have any note here")
		return errors.New("Empty")
	}

	fmt.Println("Here are all notes: ")
	fmt.Println("+------------------+")
	for _, txt := range txts {
		fmt.Println("    ", txt)
	}
	fmt.Println("+------------------+")

	return nil
}

func inputFileName() (string, error) {
	inputFileName := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name: ")
	filename, _ := inputFileName.ReadString('\n')

	finalfilename, err := trimEndLineSequence(filename)
	if len(finalfilename) < 1 || err != nil {
		return "", errors.New("Please input correct file name")
	}
	return finalfilename + ".txt", nil
}

func trimEndLineSequence(sq string) (string, error) {
	if strings.Contains(sq, "\r\n") {
		sq = strings.Replace(sq, "\r\n", "", -1)
		return sq, nil
	} else if strings.Contains(sq, "\n") {
		sq = strings.Replace(sq, "\n", "", -1)
		return sq, nil
	}
	return "", fmt.Errorf("Error while trimming end of line sequence")
}

func GetUserName() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return user.Name
}

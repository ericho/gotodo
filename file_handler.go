package main

import (
	"bufio"
	"os"
	"path"
)

const (
	TODOFILE   = ".gotodo.txt"
	FILEHEADER = "GOTODO TASKS\n========\n"
)

func getFileName() (filename string) {
	// Get env variable only works on Linux
	homefolder := os.Getenv("HOME")
	filepath := path.Join(homefolder, TODOFILE)
	return filepath
}

func readFileIntoArray(filename string) (lines []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}


func fileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func createNewFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return err
}

func writeStringToFile(filename, text string) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
	return err
}

func AddStringToFile(text string) (err error) {
	filename := getFileName()
	return writeStringToFile(filename, text)
}

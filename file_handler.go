package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

const (
	TODOFILE   = ".gotodo.txt"
	FILEHEADER = "GOTODO\n========\n"
)

// Check if file ~/.gotodo exists
// Create a new one if doesn't exists
func InitTaskFile() (err error) {
	// TODO: Fix to get $HOME environment variable
	filepath := path.Join("/home/erich", TODOFILE)
	if !fileExist(filepath) {
		err = createNewFile(filepath)
	}

	err = writeStringToFile(filepath, FILEHEADER)
	return err
}

func AddTaskToFile(task string) error {
	if len(task) == 0 {
		return fmt.Errorf("Empty tasks provided.")
	}
	fmt.Printf("Adding the task...")
	return nil
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

func MarkTaskAsDone(id int) error {
	return nil
}

func ListAllTasks() error {
	return nil
}

func fileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func createNewFile(filename string) error {
	//f, err := os.Create(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return err
}

func writeStringToFile(filename, text string) (err error) {
	// Ugly file opening, don't like the 0666..
	f, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
	return err
}

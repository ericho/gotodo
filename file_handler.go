package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
)

const (
	TODOFILE   = ".gotodo.txt"
	FILEHEADER = "GOTODO\n========\n"
)

func getFileName() (filename string) {
	// Get env variable only works on Linux
	homefolder := os.Getenv("HOME")
	filepath := path.Join(homefolder, TODOFILE)
	return filepath
}
// Check if file ~/.gotodo exists
// Create a new one if doesn't exists
func InitTaskFile() (err error) {
	filepath := getFileName()
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

func parseIdFromTaskLine(line string) (id int64, err error) {
	re := regexp.MustCompile("^[0-9]+")
	var number = re.FindString(line)
	id, err = strconv.ParseInt(number, 10, 64)
	return id, err
}

func getLastId() (id int64){
	filename := getFileName()
	lines, err := readFileIntoArray(filename)
	if err != nil {
		panic(err)
	}
	var lastId int64
	for _, i := range lines {
		lastId, _ = parseIdFromTaskLine(i)
	}
	return lastId
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

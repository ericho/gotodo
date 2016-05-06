package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func InitTaskFile() (err error) {
	filepath := getFileName()
	if !fileExist(filepath) {
		err = createNewFile(filepath)
		if err != nil {
			panic(err)
		}
		err = writeStringToFile(filepath, FILEHEADER)
	}
	return err
}

func AddTaskToFile(task string) (err error) {
	if len(task) == 0 {
		return fmt.Errorf("Empty tasks provided.")
	}
	id := getLastId()
	task = formTaskString(task, id + 1)
	fmt.Printf("Adding the task... %s\n", task)
	err = AddStringToFile(task)
	return err
}

func MarkTaskAsDone(id int) error {
	return nil
}

func ListAllTasks() error {
	// Open the file and print all it's content
	filename := getFileName()
	lines, err := readFileIntoArray(filename)
	if err != nil {
		panic(err)
	}
	for _, i := range lines {
		fmt.Printf("%s\n", i)
	}
	return err
}

func formTaskString(input string, id int64) (string) {
	return fmt.Sprintf("%d     %s\n", id, input)
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


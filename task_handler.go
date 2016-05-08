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
	task = formTaskString(task, id+1)
	err = AddStringToFile(task)
	return err
}

func MarkTaskAsDone(id_str string) error {
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		panic(err)
	}

	filename := getFileName()
	lines, err := readFileIntoArray(filename)
	if err != nil {
		panic(err)
	}

	var r int64
	for idx, i := range lines {
		if r, _ = parseIdFromTaskLine(i); id == r {
			lines[idx], _ = markLineAsDone(i)
		} else {
			lines[idx] = fmt.Sprintf("%s\n", i)
		}
	}

	err = updateStringInFile(lines)
	return err
}

func markLineAsDone(line string) (str string, err error) {
	valid := isValidLine(line)
	if !valid {
		return line, fmt.Errorf("Line %s doesn't match regexp.", line)
	}

	if done := isDoneLine(line); done {
		return line, nil
	}

	line = insertDoneMark(line)

	return line, nil
}

func isValidLine(line string) bool {
	re := regexp.MustCompile(`^[0-9]+\s+\S+`)
	return re.MatchString(line)
}

func isDoneLine(line string) bool {
	re := regexp.MustCompile(`^[0-9]+\s+\[x\].*`)
	return re.MatchString(line)
}

func insertDoneMark(line string) string {
	if isValidLine(line) {
		id, _ := parseIdFromTaskLine(line)
		mark := "[x]"
		task := parseTaskTextFromLine(line)
		line = fmt.Sprintf("%d %s %s\n", id, mark, task)
	}
	return line

}

func parseTaskTextFromLine(line string) (task string) {
	re := regexp.MustCompile(`(^[0-9]+\s+\[x\]|^[0-9]+\s+)`)
	idx := re.FindStringIndex(line)
	return line[idx[1]:]
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

func formTaskString(input string, id int64) string {
	return fmt.Sprintf("%d     %s\n", id, input)
}

func parseIdFromTaskLine(line string) (id int64, err error) {
	re := regexp.MustCompile("^[0-9]+")
	var number = re.FindString(line)
	id, err = strconv.ParseInt(number, 10, 64)
	return id, err
}

func getLastId() (id int64) {
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

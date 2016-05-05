package main

import "testing"

func TestAddTaskToFile(t *testing.T) {
	s := ""
	err := AddTaskToFile(s)
	if err == nil {
		t.Fatal("Error expected not found")
	}
}

func TestFileExistInvalidfile(t *testing.T) {
	filename := "/some/incorrect/path.txt"
	if exist := fileExist(filename); exist {
		t.Fatal("Incorrect file exists.")
	}
}

func TestFileExistValidFile(t *testing.T) {
	filename := "/proc/cpuinfo"
	if exist := fileExist(filename); !exist {
		t.Fatal("Correct file doesn't exist.")
	}
}

func TestReadFileIntoArray(t *testing.T) {
	filename := "/proc/cpuinfo"
	if l, _ := readFileIntoArray(filename); len(l) == 0 {
		t.Fatal("File cannot be read")
	}
}

func TestReadFileIntoArrayInvalidArgs(t *testing.T) {
	filename := "Somefile.txt"
	if l, err := readFileIntoArray(filename); len(l) != 0 && err != nil {
		t.Fatal("File cannot be read")
	}
}

func TestParseIdFromLine(t *testing.T) {
	id, err := parseIdFromTaskLine("10 [x] Some task")
	if err != nil {
		t.Errorf("Some error parsing regexp")
	}
	if id != 10 {
		t.Errorf("Incorrect id number")
	}
}

func TestParseIdFromLine1000(t *testing.T) {
	id, err := parseIdFromTaskLine("1000 [x] Some task")
	if err != nil {
		t.Errorf("Some error parsing regexp")
	}
	if id != 1000 {
		t.Errorf("Incorrect id number")
	}
}

func TestParseIdFromLineEmptyID(t *testing.T) {
	id, err := parseIdFromTaskLine("[x] Some task")
	if err == nil {
		t.Errorf("Some error parsing regexp")
	}
	if id != 0 {
		t.Errorf("Incorrect id number %d", id)
	}
}

func TestGetLastId(t *testing.T) {
	id := getLastId()
	if id == 0 {
		t.Errorf("Bad last Id %d", id)
	}
}
